package entity

import "github.com/peifengll/task-rebot/domain/repository"

// 订阅服务
// 用户与服务的对应关系
type Subscribe struct {
	Id        uint
	AuId      string `gorm:"column:auid"`
	ServiceId int    `gorm:"column:serviceid"`
}

type SubscribeAgg struct {
	*Subscribe
	Repo repository.SubscribeRepo
}
