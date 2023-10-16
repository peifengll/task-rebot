package test

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
)

// 热登陆存储接口
//type HotReloadStorage io.ReadWriter

func main() {
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
	//查询公众号
	mps, err := self.Friends() // self.Mps(true)
	fmt.Printf("%+v", mps)
	// 找到一个好友
	fs, err := self.Friends()
	results := fs.SearchByRemarkName(1, "刘俊杰")
	//text, err := results[0].SendText("测试消息")
	userpr := func(friend *openwechat.Friend) {
		fmt.Println(friend.User)
		fmt.Println(friend.UserName)
		fmt.Println(friend.NickName)
		fmt.Println(friend.VerifyFlag)
		fmt.Println("id?", friend.Uin)

	}
	userpr(results[0])
	results = fs.SearchByRemarkName(1, "刘俊杰")

	//if err != nil {
	//	fmt.Println(text)
	//	fmt.Println("发送失败", err)
	//}
	if err != nil {
		return
	}
	fmt.Println("结果如下：", results)
	err = bot.Block()
	if err != nil {
		fmt.Println("err is ", err)
		return
	}
}
