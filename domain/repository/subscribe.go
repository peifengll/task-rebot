package repository

import (
	"context"
	"github.com/peifengll/task-rebot/domain/entity"
)

type SubscribeRepo interface {
	//	按照服务id 查询信息，
	FindUserInfo(ctx context.Context, serviceId int) ([]*string, error)
	// 添加一个
	AddOneRecord(ctx context.Context, subscribe *entity.Subscribe) error
}
