package main

import (
	log "github.com/cihub/seelog"
	"github.com/peifengll/task-rebot/infra/component/config"
)

func Init() {
	config.InitLoadLog()
	log.Info("项目启动")
	config.InitConfig()
}

func main() {
	defer log.Flush()
	defer func() {
		log.Info("项目关闭")
	}()
	Init()

}
