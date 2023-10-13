package wechatBot

import (
	log "github.com/cihub/seelog"
	"github.com/eatmoreapple/openwechat"
)

type WechatBot struct {
	*openwechat.Bot
}

// 获取所有好友的信息
func (b *WechatBot) GetAllFriends() (openwechat.Friends, error) {
	user, err := b.GetCurrentUser()
	if err != nil {
		log.Errorf("error getting current user: %v", err)
		return nil, err
	}
	friends, err := user.Friends()
	if err != nil {
		log.Errorf("error getting friends: %s ", err)
		return nil, err
	}
	return friends, nil
}

// 这个auid就是image 生成的id
func (b *WechatBot) GetFriendByAuId(id string) (*openwechat.Friend, error) {
	Friends, _ := b.GetAllFriends()
	one := Friends.GetByRemarkName(id)
	return one, nil
}

func (b *WechatBot) GiveMessageToOne(one *openwechat.Friend, message string) error {
	_, err := one.SendText(message)
	if err != nil {
		log.Errorf("error sending message to one: %s", err)
		return err
	}
	return nil
}

// 注册消息处理逻辑
func (b *WechatBot) SetMessageSolveLogic() {
	b.Bot.MessageHandler = func(msg *openwechat.Message) {
		HandleMsg(msg)
	}

}

// 初始化一个机器人
func NewWechatBot() *WechatBot {
	// 桌面模式
	bot := openwechat.DefaultBot(openwechat.Desktop)
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl
	// 创建热存储容器对象
	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")
	defer reloadStorage.Close()
	//bot.MessageHandler = func(msg *openwechat.Message) {
	//	log.Debugf("？？？？、")
	//	HandleMsg(msg)
	//}
	err := bot.HotLogin(reloadStorage, openwechat.NewRetryLoginOption())
	if err != nil {
		log.Errorf("机器人登录失败 %s", err)
		return nil
	}
	// 获取登录后的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		log.Errorf("获取登录后的用户失败 %s", err)
	}
	log.Infof("登录成功,当前机器人为 %s", self)
	return &WechatBot{
		bot,
	}
}
