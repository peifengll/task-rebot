package repository

import (
	"context"
	"github.com/peifengll/task-rebot/domain/entity"
)

type ServiceRepo interface {
	// 建立一个新的服务
	CreateService(ctx context.Context, service *entity.Service) error
	// 查询一个服务
	FindOneService(ctx context.Context, id int) (*entity.Service, error)
	//	 查询所有启用的服务
	FindAllServicesOn(ctx context.Context) ([]*entity.Service, error)
	//	 查询所有服务
	FindAllServices(ctx context.Context) ([]*entity.Service, error)
}
