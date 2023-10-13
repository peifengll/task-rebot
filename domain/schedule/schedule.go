package schedule

import (
	"gorm.io/gorm"
	"time"
)

// 太长时间的就放入数据库？
type schedule struct {
	gorm.Model
	auid       string    `gorm:"column:auid"`
	noticetime time.Time `gorm:"column:noticetime"` // 这是几点去通知它
	number     int       `gorm:"column:number"`
	status     int       `gorm:"column:status"` //有无被通知过
	time       int       `gorm:"column:time"`   //通知几次
}

func (s *schedule) TableName() string {
	return "schedule"
}
