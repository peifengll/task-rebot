package entity

type Service struct {
	Id         uint
	Name       string `gorm:"column:name"`
	ServerTime string `gorm:"column:servertime"` //服务上线的时间
	Status     int    `gorm:"column:status"`     //是否启用
}
