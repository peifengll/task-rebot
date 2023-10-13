package wechatBot

import (
	"context"
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/eatmoreapple/openwechat"
	"github.com/peifengll/task-rebot/app"
	"github.com/peifengll/task-rebot/domain/task"
	"strconv"
	"strings"
)

const (
	// 帮助
	messageType3 = "-c"
	// 提醒
	messageType1 = "+s"
	// 任务
	messageType2 = "+t"
	// 一堆提醒在redis存哪个键
	messageType1Key = "weixin:schedule.notice"
	// 一堆任务在redis存哪个键
	messageType2Key = "weixin:task.notice"
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
		// 掐掉头尾的空格
		contentText = strings.TrimSpace(msg.Content)
		SolveTextMessage(msg, contentText, sender.ID())
	}

	// 最后一定点一个已读
	msg.AsRead()
}

// // 返回是什么指令
//
//	func HasString(text string) string {
//		if strings.HasSuffix(text, messageType1) {
//			return messageType1
//		}
//		if strings.HasSuffix(text, messageType2) {
//			return messageType2
//		}
//		return ""
//	}
func SolveTextMessage(msg *openwechat.Message, text, userid string) {
	if text == "帮助" {
		helpinfo := `
本人提供功能如下:
// 注:本人将用户的头像作为用户的唯一标识，如若更换头像，那么以前的信息将对你失效

【任务管理】
	+t20220101 11:11,任务名称,任务内容  //会在2022.01.01 11:11之前发消息(该消息会包含任务id)进行提醒，若逾期该任务会被计为失败
	完成task_$id      // 代表编号为$id的任务已完成，会将其进行记录，并不会再进行提醒
	今日情况 		// 展示截止日期为今日的所有任务的情况
【定时提醒】
	+s20220101 11:11,消息内容  //会在 2022.01.01 11:11 时向你发送消息要提醒的内容
`
		msg.ReplyText(helpinfo)
		return
	}

	if text == "今日情况" {
		//	todo 展示截止日期为今日的所有任务的情况
		ctx := context.Background()
		tasks, err := app.TaskRepo.FindDueTaskPerson(ctx, userid, -1)
		if err != nil {
			log.Errorf("TaskRepo.FindDueTaskPerson failed with error: %v", err)
			msg.ReplyText("出错了、请联系管理员")
			return
		}
		total := len(tasks)
		st := "今日总计有" + strconv.Itoa(total) + "个任务\n"
		a := make([]*task.Task, 0)
		b := make([]*task.Task, 0)
		for i := 0; i < total; i++ {
			if tasks[i].Status == 0 {
				a = append(a, tasks[i])
			} else if tasks[i].Status == 1 {
				b = append(b, tasks[i])
			}
		}
		st += "其中," + strconv.Itoa(len(b)) + "件任务已完成,分别为:\n"
		st += "----"
		for i := 0; i < len(b); i++ {
			st += fmt.Sprintf("第%d:\n 任务ID: %d\n 任务名称: %s\n,任务内容: %s \n====\n", i+1, b[i].Id, b[i].Name, b[i].Description)
		}
		st += strconv.Itoa(len(a)) + "件任务未完成,分别为:\n"
		for i := 0; i < len(a); i++ {
			st += fmt.Sprintf("第%d:\n  任务ID: %d\n 任务名称: %s\n,任务内容: %s \n====\n", i+1, a[i].Id, a[i].Name, a[i].Description)
		}
		st += "\n"
		msg.ReplyText(st)
		return
	}

	if strings.HasPrefix(text, "完成task_") {
		id := strings.TrimPrefix(text, "完成task_")
		iid, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			log.Errorf("strconv.ParseInt failed with error: %v, id: %d", err, iid)
			msg.ReplyText("请检查输入,task_之后应该为一个整数id")
			return
		}

		//
		//	 todo 去更新这个task，同时删除redis中的  redis还没写呢
		return
	}

	if strings.HasPrefix(text, "+s") {
		// todo  +s     这个也要等等  还是新建一个表吧.
		return
	}

	if strings.HasPrefix(text, "+t") {
		//	todo +t
		// 处理text
		return
	}

	msg.ReplyText("??????")
	log.Infof("msg.IsText excute?")
}

func SolveNewFriend(msg *openwechat.Message) {
	friend, err := msg.Agree()
	if err != nil {
		log.Errorf("error agreeing friend: %v", err)
	}
	_, err = friend.SendText("你好,我是你的小助手,你可以通过发送\"帮助\"来获取帮助信息")
	if err != nil {
		log.Errorf("error send  tips: %v", err)
		return
	}
}
