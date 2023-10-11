package main

import (
	log "github.com/cihub/seelog"
	"github.com/peifengll/task-rebot/app"
	"github.com/peifengll/task-rebot/domain/wechatBot"
	"github.com/peifengll/task-rebot/infra/component/config"
)

func main() {
	config.InitLoadLog()
	config.InitConfig()
	defer log.Flush()
	app.Init()

	bot := wechatBot.NewWechatBot()
	bot.SetMessageSolveLogic()
	bot.Block()
}
