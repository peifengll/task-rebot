package task

import (
	"time"
)

type Task struct {
	Id          uint
	AuId        string    `gorm:"column:auid"`        // 属于哪个用户
	Name        string    `gorm:"column:name"`        //任务的名字
	Description string    `gorm:"column:description"` //细节
	DueDate     time.Time `gorm:"column:duedate"`     //截至日期
	Status      int       `gorm:"column:status"`      // 情况，0没弄，1完成，2失败超时

}

type TaskAgg struct {
	*Task
	Repo TaskRepo
}

// 任务管理者
