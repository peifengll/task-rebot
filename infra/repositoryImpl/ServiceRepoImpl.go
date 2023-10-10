package repositoryImpl

import (
	"context"
	log "github.com/cihub/seelog"
	"github.com/peifengll/task-rebot/domain/entity"
	"github.com/peifengll/task-rebot/domain/repository"
	"github.com/peifengll/task-rebot/infra/converter"
	"github.com/peifengll/task-rebot/infra/persistence/dao"
)

type ServiceRepoImpl struct {
	Dao     dao.ServiceDao
	Convert converter.ServiceConverter
}

var _ repository.ServiceRepo = &ServiceRepoImpl{}

func NewServiceRepo(dao dao.ServiceDao, convert converter.ServiceConverter) repository.ServiceRepo {
	return &ServiceRepoImpl{
		Dao:     dao,
		Convert: convert,
	}
}

// 建立一个新的服务
func (s *ServiceRepoImpl) CreateService(ctx context.Context, service *entity.Service) error {
	po := s.Convert.ToPoService(service)
	err := s.Dao.CreateService(ctx, po)
	if err != nil {
		log.Errorf("Failed to create service %s, error: %v", service.Name, err)
		return err
	}
	return nil
}

// 查询一个服务
func (s *ServiceRepoImpl) FindOneService(ctx context.Context, id int) (*entity.Service, error) {
	po, err := s.Dao.FindServiceById(ctx, id)
	if err != nil {
		log.Errorf("Failed to find service %d, error: %v", id, err)
		return nil, err
	}
	return s.Convert.ToEntityService(po), nil
}

// 查询所有启用的服务
func (s *ServiceRepoImpl) FindAllServicesOn(ctx context.Context) ([]*entity.Service, error) {
	poList, err := s.Dao.FindAllServiceStatus(ctx, 1)
	if err != nil {
		log.Errorf("Failed to find all services on, error: %v", err)
		return nil, err
	}
	return s.Convert.ToEntityServices(poList), nil
}

// 查询所有服务
func (s *ServiceRepoImpl) FindAllServices(ctx context.Context) ([]*entity.Service, error) {
	poList, err := s.Dao.FindAllService(ctx)
	if err != nil {
		log.Errorf("Failed to find all services, error: %v", err)
		return nil, err
	}
	return s.Convert.ToEntityServices(poList), nil
}
