package config

import (
	"fmt"
	"github.com/peifengll/task-rebot/domain/entity"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Init() {
	cfg, err := ini.Load("../conf.ini")
	if err != nil {
		log.Fatal("Fail to read file: ", err)
	}

	fmt.Println("App Name:", cfg.Section("").Key("app_name").String())
	fmt.Println("Log Level:", cfg.Section("").Key("log_level").String())

	fmt.Println("MySQL IP:", cfg.Section("mysql").Key("ip").String())
	mysqlPort, err := cfg.Section("mysql").Key("port").Int()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MySQL Port:", mysqlPort)
	fmt.Println("MySQL User:", cfg.Section("mysql").Key("user").String())
	fmt.Println("MySQL Password:", cfg.Section("mysql").Key("password").String())
	fmt.Println("MySQL Database:", cfg.Section("mysql").Key("database").String())

	fmt.Println("Redis IP:", cfg.Section("redis").Key("ip").String())
	redisPort, err := cfg.Section("redis").Key("port").Int()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Redis Port:", redisPort)
}

func InitDb() {
	dsn := "root:123456@(118.25.23.154:8086)/taskrebot?charset=utf8&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	DB = db
	db.AutoMigrate(&entity.User{}, &entity.Task{})
}

// 抛出DB
func GetDb() *gorm.DB {
	if DB != nil {
		return DB
	} else {
		return nil
	}
}
