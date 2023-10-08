package dao

import (
	"context"
	"github.com/peifengll/task-rebot/infra/persistence/po"
)

// 大体就是服务的增删查改功能
type ServiceDao interface {
	//按照id查询service
	FindServiceById(ctx context.Context, id int) (*po.Service, error)
	//查询所有服务
	FindAllService(ctx context.Context) ([]*po.Service, error)
	//	按状态查询服务
	FindAllServiceStatus(ctx context.Context, status int) ([]*po.Service, error)
	//	按照该服务启动的时间查询服务
	FindServicesByServerTime(ctx context.Context, servertime string) ([]*po.Service, error)
	//	添加服务记录
	CreateService(ctx context.Context, services *po.Service) error
	//修改服务是否启用
	UpdateServiceStatus(ctx context.Context, id int, status int) error
	//	删除一个服务
	DeleteService(ctx context.Context, id int) error
}
