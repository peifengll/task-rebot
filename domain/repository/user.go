package repository

import (
	"context"
	"github.com/peifengll/task-rebot/domain/entity"
)

// 存进去大概就是用不上了
type UserRepo interface {
	SaveUser(ctx context.Context, task *entity.User) error
	// 获取一个新的id
	GetNewAuId(ctx context.Context) (int, error)
}
