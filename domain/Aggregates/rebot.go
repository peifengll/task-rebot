package Aggregates

import "github.com/peifengll/task-rebot/domain/entity"

var R Robot

// 这玩意儿怎么也该算一个聚合
type Robot struct {
	*entity.ServiceAgg
	*entity.SubscribeAgg
	*entity.UserAgg
	*entity.TaskAgg
}

// 打开定时服务
