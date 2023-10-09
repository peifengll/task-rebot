package config

import (
	log "github.com/cihub/seelog"
	"github.com/spf13/viper"
)

type TomConfig struct {
	AppName string
	Mysql   MySQLConfig
}

type MySQLConfig struct {
	Host        string
	DbName      string
	Password    string
	Port        int
	TablePrefix string
	User        string
}

var c *TomConfig

//func init() {
//	fmt.Println("Init Starting")
//	InitConfig()
//	fmt.Println("Init ending")
//
//}

func InitLoadLog() {
	logger, err := log.LoggerFromConfigAsFile("infra/component/config/seelog.xml")
	if err != nil {
		log.Errorf("parse config.xml error: %v", err)
		return
	}

	log.ReplaceLogger(logger)
	//log.Error("meiwneti")
}

func InitConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("infra/component/config")
	viper.AddConfigPath("../infra/component/config") // 这个是测试路径
	err := viper.ReadInConfig()
	if nil != err {
		log.Errorf("%s", err)
		return err
	}
	err = viper.Unmarshal(&c)
	if err != nil {
		log.Errorf("%s", err)
		return err
	}
	log.Info("数据库连接配置正常")
	return nil
}
func GetConfig() (*TomConfig, error) {
	if c == nil {
		log.Info("config在GetConfig中初始化")
		err := InitConfig()
		if err != nil {
			log.Errorf("配置初始化失败: %s", err)
			return nil, err
		}
	}
	return c, nil
}
