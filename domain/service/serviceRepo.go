package service

import (
	"context"
)

type ServiceRepo interface {
	// 建立一个新的服务
	CreateService(ctx context.Context, service *Service) error
	// 查询一个服务
	FindOneService(ctx context.Context, id int) (*Service, error)
	//	 查询所有启用的服务
	FindAllServicesOn(ctx context.Context) ([]*Service, error)
	//	 查询所有服务
	FindAllServices(ctx context.Context) ([]*Service, error)
}
