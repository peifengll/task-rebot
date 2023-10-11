package converter

import (
	"github.com/peifengll/task-rebot/domain/subscribe"
	"github.com/peifengll/task-rebot/infra/persistence/po"
	"gorm.io/gorm"
)

type SubscribeConverter interface {
	ToEntitySubscribe(subscribe *po.Subscribe) *subscribe.Subscribe
	ToEntitySubscribes(subscribes []*po.Subscribe) []*subscribe.Subscribe
	TpPoSubscribe(subscribe *subscribe.Subscribe) *po.Subscribe
	ToPoSubscribes(subscribes []*subscribe.Subscribe) []*po.Subscribe
}

type SubscribeConverterImpl struct {
}

var _ SubscribeConverter = &SubscribeConverterImpl{}

func NewSubscribeConverter() SubscribeConverter {
	return &SubscribeConverterImpl{}
}

func (p *SubscribeConverterImpl) ToEntitySubscribe(s *po.Subscribe) *subscribe.Subscribe {
	return &subscribe.Subscribe{
		Id:        s.ID,
		AuId:      s.AuId,
		ServiceId: s.ServiceId,
	}
}
func (p *SubscribeConverterImpl) ToEntitySubscribes(subscribes []*po.Subscribe) []*subscribe.Subscribe {
	sublist := make([]*subscribe.Subscribe, len(subscribes))
	for i, s := range subscribes {
		sublist[i] = &subscribe.Subscribe{
			Id:        s.ID,
			AuId:      s.AuId,
			ServiceId: s.ServiceId,
		}
	}
	return sublist
}
func (p *SubscribeConverterImpl) TpPoSubscribe(subscribe *subscribe.Subscribe) *po.Subscribe {
	return &po.Subscribe{
		Model: gorm.Model{
			ID: subscribe.Id,
		},
		AuId:      subscribe.AuId,
		ServiceId: subscribe.ServiceId,
	}
}
func (p *SubscribeConverterImpl) ToPoSubscribes(subscribes []*subscribe.Subscribe) []*po.Subscribe {
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
