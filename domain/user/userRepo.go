package user

import (
	"context"
)

// 存进去大概就是用不上了
type UserRepo interface {
	SaveUser(ctx context.Context, user *User) error
	// 获取一个新的id
	GetNewAuId(ctx context.Context) (int, error)
}
