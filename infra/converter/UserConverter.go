package converter

import (
	"github.com/peifengll/task-rebot/domain/user"
	"github.com/peifengll/task-rebot/infra/persistence/po"
)

type UserConverter interface {
	ToEntityUser(user *po.User) *user.User
	ToEntityUsers(users []*po.User) []*user.User
	ToPoUser(user *user.User) *po.User
	ToPoUsers(users []*user.User) []*po.User
}

type UserConverterImpl struct {
}

func NewUserConverter() UserConverter {
	return &UserConverterImpl{}
}

var _ UserConverter = &UserConverterImpl{}

func (r *UserConverterImpl) ToEntityUser(u *po.User) *user.User {
	return &user.User{
		AuId:     u.AuId,
		UserName: u.UserName,
	}
}
func (r *UserConverterImpl) ToEntityUsers(users []*po.User) []*user.User {
	list := make([]*user.User, len(users))
	for i, u := range users {
		list[i] = &user.User{
			AuId:     u.AuId,
			UserName: u.UserName,
		}
	}
	return list
}
func (r *UserConverterImpl) ToPoUser(user *user.User) *po.User {
	return &po.User{
		AuId:     user.AuId,
		UserName: user.UserName,
	}
}
func (r *UserConverterImpl) ToPoUsers(users []*user.User) []*po.User {
	list := make([]*po.User, len(users))
	for i, user := range users {
		list[i] = &po.User{
			AuId:     user.AuId,
			UserName: user.UserName,
		}
	}
	return list
}
