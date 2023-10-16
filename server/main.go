package main

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"github.com/peifengll/task-rebot/config"
)

var Bot *openwechat.Bot

func initBot() {
	// 桌面模式
	bot := openwechat.DefaultBot(openwechat.Desktop)
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl
	// 创建热存储容器对象
	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")
	defer reloadStorage.Close()
	err := bot.HotLogin(reloadStorage, openwechat.NewRetryLoginOption())
	//err := bot.Login()
	if err != nil {
		fmt.Println("登录失败: ", err)
		return
	}
	// 获取登录后的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		fmt.Println("获取登录后的用户失败")
	}
	fmt.Println(self)
	Bot = bot
}
func main() {
	// 完成初始化
	config.InitDb() //  初始化数据库
	initBot()       // 机器人，初始化
	//	开出子协程

	Bot.Block()
}
