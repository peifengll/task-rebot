package dao

import (
	"context"
	"github.com/peifengll/task-rebot/infra/persistence/po"
)

type SubscribeDao interface {
	// 增加一条订阅
	CreateSubscribe(ctx context.Context, s *po.Subscribe) error
	//	删除一条订阅
	DeleteSubscribe(ctx context.Context, id int) error
	//	按找服务id查询一堆订阅,查询出用户id
	GetSubscribeByServiceId(ctx context.Context, seid int) ([]*string, error)
}
