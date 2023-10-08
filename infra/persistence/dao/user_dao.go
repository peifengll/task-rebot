package dao

import (
	"context"
	"github.com/peifengll/task-rebot/infra/persistence/po"
)

type UserDao interface {
	// 返回一个id
	CreateUser(ctx context.Context, user *po.User) error

	//	找到新的编号
	FindNewId(ctx context.Context) (int, error)
}
