package po

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	Name       string `gorm:"column:name"`
	ServerTime string `gorm:"column:servertime"` //服务上线的时间
	Status     int    `gorm:"column:status"`     //是否启用
}

func (r *Service) TableName() string {
	return "Services"
}
