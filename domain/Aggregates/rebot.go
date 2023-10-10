package Aggregates

import (
	"github.com/peifengll/task-rebot/domain/service"
	"github.com/peifengll/task-rebot/domain/subscribe"
	"github.com/peifengll/task-rebot/domain/task"
	"github.com/peifengll/task-rebot/domain/user"
)

var R Robot

// 这玩意儿怎么也该算一个聚合
type Robot struct {
	*service.ServiceAgg
	*subscribe.SubscribeAgg
	*user.UserAgg
	*task.TaskAgg
}

// 打开定时服务
