package converter

import (
	"github.com/peifengll/task-rebot/domain/entity"
	"github.com/peifengll/task-rebot/infra/persistence/po"
)

type ServiceConverter interface {
	ToEntityService(service *po.Service) *entity.Service
	ToEntityServices(services []*po.Service) []*entity.Service
	ToPoService(service *entity.Service) *po.Service
	ToPoServices(services []*entity.Service) []*po.Service
}

type ServiceConverterImpl struct {
}

func (r *ServiceConverterImpl) ToEntityService(service *po.Service) *entity.Service {
	panic("")
}
func (r *ServiceConverterImpl) ToEntityServices(services []*po.Service) []*entity.Service {
	panic("")
}
func (r *ServiceConverterImpl) ToPoService(service *entity.Service) *po.Service {
	panic("")
}
func (r *ServiceConverterImpl) ToPoServices(services []*entity.Service) []*po.Service {
	panic("")
}
