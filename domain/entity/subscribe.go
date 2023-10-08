package entity

// 订阅服务
// 用户与服务的对应关系
type Subscribe struct {
	Id        uint
	AuId      string `gorm:"column:auid"`
	ServiceId int    `gorm:"column:serviceid"`
}
