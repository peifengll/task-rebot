package converter

import (
	"github.com/peifengll/task-rebot/domain/entity"
	"github.com/peifengll/task-rebot/infra/persistence/po"
	"gorm.io/gorm"
)

type ServiceConverter interface {
	ToEntityService(service *po.Service) *entity.Service
	ToEntityServices(services []*po.Service) []*entity.Service
	ToPoService(service *entity.Service) *po.Service
	ToPoServices(services []*entity.Service) []*po.Service
}

type ServiceConverterImpl struct {
}

var _ ServiceConverter = &ServiceConverterImpl{}

func NewServiceConverter() ServiceConverter {
	return &ServiceConverterImpl{}
}

func (r *ServiceConverterImpl) ToEntityService(service *po.Service) *entity.Service {
	aim := &entity.Service{
		Id:         service.ID,
		Name:       service.Name,
		ServerTime: service.ServerTime,
		Status:     service.Status,
	}
	return aim
}
func (r *ServiceConverterImpl) ToEntityServices(services []*po.Service) []*entity.Service {
	enlist := make([]*entity.Service, len(services))
	for i, s := range services {
		enlist[i] = &entity.Service{
			Id:         s.ID,
			Name:       s.Name,
			ServerTime: s.ServerTime,
			Status:     s.Status,
		}
	}
	return enlist
}
func (r *ServiceConverterImpl) ToPoService(service *entity.Service) *po.Service {
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
func (r *ServiceConverterImpl) ToPoServices(services []*entity.Service) []*po.Service {
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
