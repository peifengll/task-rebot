package config

import (
	"fmt"
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

func init() {
	fmt.Println("Init Starting")
	InitConfig()
	fmt.Println("Init ending")

}

func InitConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("infra/component/config")
	viper.AddConfigPath("../infra/component/config") // 这个是测试路径
	err := viper.ReadInConfig()
	if nil != err {
		return err
	}
	err = viper.Unmarshal(&c)
	if err != nil {
		return err
	}

	return nil
}
func GetConfig() (*TomConfig, error) {
	if c == nil {
		fmt.Println("config在GetConfig中初始化")
		err := InitConfig()
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}
