package ingorm

import (
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/peifengll/task-rebot/infra/component/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var db *gorm.DB
var once sync.Once

func ProviderDsn(config config.MySQLConfig) string {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=10s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DbName,
	)
	log.Infof("gorm dsn: %s", dsn)
	return dsn
}

func ProviderOnceGormDB() *gorm.DB {
	if db == nil {
		once.Do(func() {
			c, err := config.GetConfig()
			if err != nil {
				log.Errorf("Error getting mysql config: %s", err)
				return
			}
			gormdb, err := gorm.Open(mysql.Open(ProviderDsn(c.Mysql)), &gorm.Config{})
			if err != nil {
				log.Error("Error opening mysql connection: %s", err)
				return
			}
			db = gormdb
		})
	}
	return db
}
