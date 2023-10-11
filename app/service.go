package app

import (
	"github.com/peifengll/task-rebot/domain/Aggregates"
	"github.com/peifengll/task-rebot/domain/service"
	"github.com/peifengll/task-rebot/domain/subscribe"
	"github.com/peifengll/task-rebot/domain/task"
	"github.com/peifengll/task-rebot/domain/user"
)

var UserRepo user.UserRepo
var TaskRepo task.TaskRepo
var ServiceRepo service.ServiceRepo
var SubscriberRepo subscribe.SubscribeRepo

func Init() {
	UserRepo = Aggregates.NewUserRepo()
	TaskRepo = Aggregates.NewTaskRepo()
	ServiceRepo = Aggregates.NewServiceRepo()
	SubscriberRepo = Aggregates.NewSubscribeRepo()
}
