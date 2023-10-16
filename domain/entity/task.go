package entity

import (
	"gorm.io/gorm"
	"time"
)

// Task 先只做一个简单的提醒功能
type Task struct {
	gorm.Model
	AuId        string    `gorm:"column:au_id"` // 属于哪个用户
	Name        string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:varchar(255);not null"`
	DueDate     time.Time `gorm:"column:due_date"`   //什么时间截止
	Status      int       `gorm:"type:int;not null"` // 状态， 0是没做，1：完成，2：超时，3：失败
}
