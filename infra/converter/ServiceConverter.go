package converter

import (
	"github.com/peifengll/task-rebot/domain/service"
	"github.com/peifengll/task-rebot/infra/persistence/po"
	"gorm.io/gorm"
)

type ServiceConverter interface {
	ToEntityService(service *po.Service) *service.Service
	ToEntityServices(services []*po.Service) []*service.Service
	ToPoService(service *service.Service) *po.Service
	ToPoServices(services []*service.Service) []*po.Service
}

type ServiceConverterImpl struct {
}

var _ ServiceConverter = &ServiceConverterImpl{}

func NewServiceConverter() ServiceConverter {
	return &ServiceConverterImpl{}
}

func (r *ServiceConverterImpl) ToEntityService(s *po.Service) *service.Service {
	aim := &service.Service{
		Id:         s.ID,
		Name:       s.Name,
		ServerTime: s.ServerTime,
		Status:     s.Status,
	}
	return aim
}
func (r *ServiceConverterImpl) ToEntityServices(services []*po.Service) []*service.Service {
	enlist := make([]*service.Service, len(services))
	for i, s := range services {
		enlist[i] = &service.Service{
			Id:         s.ID,
			Name:       s.Name,
			ServerTime: s.ServerTime,
			Status:     s.Status,
		}
	}
	return enlist
}
func (r *ServiceConverterImpl) ToPoService(service *service.Service) *po.Service {
	aim := &po.Service{
		Model: gorm.Model{
			ID: service.Id,
		},
		Name:       service.Name,
		ServerTime: service.ServerTime,
		Status:     service.Status,
	}
	return aim
}
func (r *ServiceConverterImpl) ToPoServices(services []*service.Service) []*po.Service {
	enlist := make([]*po.Service, len(services))
	for i, s := range services {
		enlist[i] = &po.Service{
			Model: gorm.Model{
				ID: s.Id,
			},
			Name:       s.Name,
			ServerTime: s.ServerTime,
			Status:     s.Status,
		}
	}
	return enlist
}
