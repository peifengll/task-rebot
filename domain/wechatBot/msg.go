package wechatBot

import (
	log "github.com/cihub/seelog"
	"github.com/eatmoreapple/openwechat"
)

func HandleMsg(msg *openwechat.Message) {
	var (
		contentText = ""
		sender      *openwechat.User
		err         error
	)
	sender, err = msg.Sender()
	if err != nil {
		log.Errorf("获取发送人消息失败")
		msg.ReplyText(err.Error())
		return
	}
	// 如果是新申请的好友
	if msg.IsFriendAdd() {
		SolveNewFriend(msg)
	} else if msg.IsText() {
		contentText = msg.Content
		SolveTextMessage(contentText, sender.ID())
	}
	// 最后一定点一个已读
	msg.AsRead()
}

func SolveTextMessage(text, userid string) {

}

func SolveNewFriend(msg *openwechat.Message) {
	friend, err := msg.Agree()
	if err != nil {
		log.Errorf("error agreeing friend: %v", err)
	}
	_, err = friend.SendText("你好,我是你的小助手,你可以通过发送'-h'来获取帮助信息")
	if err != nil {
		log.Errorf("error send  tips: %v", err)
		return
	}
}
