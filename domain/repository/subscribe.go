package repository

import (
	"context"
	"github.com/peifengll/task-rebot/domain/entity"
)

type SubscribeRepo interface {
	//	按照服务id 查询信息，
	FindInfo(ctx context.Context, serviceId int) ([]*entity.Subscribe, error)
	AddOneRecord(ctx context.Context, subscribe *entity.Subscribe) error
}
