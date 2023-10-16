package entity

import "gorm.io/gorm"

// User user其实都是这个账号所对应的好友，
type User struct {
	gorm.Model
	AuId     string `gorm:"column:au_id"`     // 给好友的备注，也是寻找好友的唯一标识
	Username string `gorm:"column:username"`  // 昵称
	NotifyOn bool   `gorm:"column:notify_on"` // 是否开启通知提醒
}
