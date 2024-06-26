package po

import "gorm.io/gorm"

// 订阅服务
// 用户与服务的对应关系
type Subscribe struct {
	gorm.Model
	AuId      string `gorm:"column:auid"`
	ServiceId int    `gorm:"column:serviceid"`
}

func (s *Subscribe) TableName() string {
	return "subscribe"
}
