package converter

import (
	"github.com/peifengll/task-rebot/domain/entity"
	"github.com/peifengll/task-rebot/infra/persistence/po"
)

type SubscribeConverter interface {
	ToEntitySubscribe(subscribe *po.Subscribe) *entity.Subscribe
	ToEntitySubscribes(subscribes []*po.Subscribe) []*entity.Subscribe
	TpPoSubscribe(subscribes *entity.Subscribe) *po.Subscribe
	ToPoSubscribes(subscribes []*entity.Subscribe) []*po.Subscribe
}
