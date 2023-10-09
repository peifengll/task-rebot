package repositoryImpl

import (
	"context"
	"github.com/peifengll/task-rebot/domain/entity"
	"github.com/peifengll/task-rebot/infra/persistence/dao"
)

type ServiceRepoImpl struct {
	Dao dao.ServiceDao
	//	todo
}

// 建立一个新的服务
func (s *ServiceRepoImpl) CreateService(ctx context.Context, service *entity.Service) error {
	panic("")
}

// 查询一个服务
func (s *ServiceRepoImpl) FindOneService(ctx context.Context, id int) (*entity.Service, error) {
	panic("")

}

// 查询所有启用的服务
func (s *ServiceRepoImpl) FindAllServicesOn(ctx context.Context) ([]*entity.Service, error) {
	panic("")
}

// 查询所有服务
func (s *ServiceRepoImpl) FindAllServices(ctx context.Context) ([]*entity.Service, error) {
	panic("")
}
