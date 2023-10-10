package service

type Service struct {
	Id         uint
	Name       string `gorm:"column:name"`
	ServerTime string `gorm:"column:servertime"` //服务上线的时间
	Status     int    `gorm:"column:status"`     //是否启用
}

type ServiceAgg struct {
	*Service
	dict map[string]ServiceVirtual
	Repo ServiceRepo
}

// 以代码形式存储的服务
type ServiceVirtual struct {
	*Service
	run func() error
}
