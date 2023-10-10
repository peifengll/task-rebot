package converter

import (
	"github.com/peifengll/task-rebot/domain/entity"
	"github.com/peifengll/task-rebot/infra/persistence/po"
	"gorm.io/gorm"
)

type SubscribeConverter interface {
	ToEntitySubscribe(subscribe *po.Subscribe) *entity.Subscribe
	ToEntitySubscribes(subscribes []*po.Subscribe) []*entity.Subscribe
	TpPoSubscribe(subscribe *entity.Subscribe) *po.Subscribe
	ToPoSubscribes(subscribes []*entity.Subscribe) []*po.Subscribe
}

type SubscribeConverterImpl struct {
}

var _ SubscribeConverter = &SubscribeConverterImpl{}

func NewSubscribeConverter() SubscribeConverter {
	return &SubscribeConverterImpl{}
}

func (p *SubscribeConverterImpl) ToEntitySubscribe(subscribe *po.Subscribe) *entity.Subscribe {
	return &entity.Subscribe{
		Id:        subscribe.ID,
		AuId:      subscribe.AuId,
		ServiceId: subscribe.ServiceId,
	}
}
func (p *SubscribeConverterImpl) ToEntitySubscribes(subscribes []*po.Subscribe) []*entity.Subscribe {
	sublist := make([]*entity.Subscribe, len(subscribes))
	for i, s := range subscribes {
		sublist[i] = &entity.Subscribe{
			Id:        s.ID,
			AuId:      s.AuId,
			ServiceId: s.ServiceId,
		}
	}
	return sublist
}
func (p *SubscribeConverterImpl) TpPoSubscribe(subscribe *entity.Subscribe) *po.Subscribe {
	return &po.Subscribe{
		Model: gorm.Model{
			ID: subscribe.Id,
		},
		AuId:      subscribe.AuId,
		ServiceId: subscribe.ServiceId,
	}
}
func (p *SubscribeConverterImpl) ToPoSubscribes(subscribes []*entity.Subscribe) []*po.Subscribe {
	sublist := make([]*po.Subscribe, len(subscribes))
	for i, s := range subscribes {
		sublist[i] = &po.Subscribe{
			Model: gorm.Model{
				ID: s.Id,
			},
			AuId:      s.AuId,
			ServiceId: s.ServiceId,
		}
	}
	return sublist
}
