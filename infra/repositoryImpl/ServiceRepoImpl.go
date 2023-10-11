package repositoryImpl

import (
	"context"
	log "github.com/cihub/seelog"
	"github.com/peifengll/task-rebot/domain/service"
	"github.com/peifengll/task-rebot/infra/converter"
	"github.com/peifengll/task-rebot/infra/persistence/dao"
)

type ServiceRepoImpl struct {
	Dao     dao.ServiceDao
	Convert converter.ServiceConverter
}

var _ service.ServiceRepo = &ServiceRepoImpl{}

func NewServiceRepo(dao dao.ServiceDao, convert converter.ServiceConverter) service.ServiceRepo {
	return &ServiceRepoImpl{
		Dao:     dao,
		Convert: convert,
	}
}

// 建立一个新的服务
func (s *ServiceRepoImpl) CreateService(ctx context.Context, service *service.Service) error {
	po := s.Convert.ToPoService(service)
	err := s.Dao.CreateService(ctx, po)
	if err != nil {
		log.Errorf("Failed to create domain_service %s, error: %v", service.Name, err)
		return err
	}
	return nil
}

// 查询一个服务
func (s *ServiceRepoImpl) FindOneService(ctx context.Context, id int) (*service.Service, error) {
	po, err := s.Dao.FindServiceById(ctx, id)
	if err != nil {
		log.Errorf("Failed to find domain_service %d, error: %v", id, err)
		return nil, err
	}
	return s.Convert.ToEntityService(po), nil
}

// 查询所有启用的服务
func (s *ServiceRepoImpl) FindAllServicesOn(ctx context.Context) ([]*service.Service, error) {
	poList, err := s.Dao.FindAllServiceStatus(ctx, 1)
	if err != nil {
		log.Errorf("Failed to find all services on, error: %v", err)
		return nil, err
	}
	return s.Convert.ToEntityServices(poList), nil
}

// 查询所有服务
func (s *ServiceRepoImpl) FindAllServices(ctx context.Context) ([]*service.Service, error) {
	poList, err := s.Dao.FindAllService(ctx)
	if err != nil {
		log.Errorf("Failed to find all services, error: %v", err)
		return nil, err
	}
	return s.Convert.ToEntityServices(poList), nil
}
