//go:build wireinject
// +build wireinject

package Aggregates

import (
	"github.com/google/wire"
	"github.com/peifengll/task-rebot/domain/service"
	"github.com/peifengll/task-rebot/domain/subscribe"
	"github.com/peifengll/task-rebot/domain/task"
	"github.com/peifengll/task-rebot/domain/user"
	"github.com/peifengll/task-rebot/infra/component/mysql/ingorm"
	"github.com/peifengll/task-rebot/infra/converter"
	"github.com/peifengll/task-rebot/infra/persistence/dao"
	"github.com/peifengll/task-rebot/infra/repositoryImpl"
)

func NewServiceRepo() service.ServiceRepo {
	wire.Build(repositoryImpl.NewServiceRepo, ingorm.ProviderOnceGormDB,
		dao.NewServiceDao, converter.NewServiceConverter,
	)
	return &repositoryImpl.ServiceRepoImpl{}
}

func NewTaskRepo() task.TaskRepo {
	wire.Build(repositoryImpl.NewTaskRepo, ingorm.ProviderOnceGormDB,
		dao.NewTaskDao, converter.NewTaskConverter,
	)
	return &repositoryImpl.TaskRepoImpl{}
}
func NewSubscribeRepo() subscribe.SubscribeRepo {
	wire.Build(repositoryImpl.NewSubscribeRepo, ingorm.ProviderOnceGormDB,
		dao.NewSubscribeDao, converter.NewSubscribeConverter)
	return &repositoryImpl.SubscribeRepoImpl{}
}
func NewUserRepo() user.UserRepo {
	wire.Build(repositoryImpl.NewUserRepo, ingorm.ProviderOnceGormDB,
		dao.NewUserDao, converter.NewUserConverter)
	return &repositoryImpl.UserRepoImpl{}
}
