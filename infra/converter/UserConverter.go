package converter

import (
	"github.com/peifengll/task-rebot/domain/entity"
	"github.com/peifengll/task-rebot/infra/persistence/po"
)

type UserConverter interface {
	ToEntityUser(user *po.User) *entity.User
	ToEntityUsers(users []*po.User) []*entity.User
	ToPoUser(user *entity.User) *po.User
	ToPoUsers(users []*entity.User) []*po.User
}

type UserConverterImpl struct {
}

func NewUserConverter() UserConverter {
	return &UserConverterImpl{}
}

var _ UserConverter = &UserConverterImpl{}

func (r *UserConverterImpl) ToEntityUser(user *po.User) *entity.User {
	return &entity.User{
		AuId:     user.AuId,
		UserName: user.UserName,
	}
}
func (r *UserConverterImpl) ToEntityUsers(users []*po.User) []*entity.User {
	list := make([]*entity.User, len(users))
	for i, user := range users {
		list[i] = &entity.User{
			AuId:     user.AuId,
			UserName: user.UserName,
		}
	}
	return list
}
func (r *UserConverterImpl) ToPoUser(user *entity.User) *po.User {
	return &po.User{
		AuId:     user.AuId,
		UserName: user.UserName,
	}
}
func (r *UserConverterImpl) ToPoUsers(users []*entity.User) []*po.User {
	list := make([]*po.User, len(users))
	for i, user := range users {
		list[i] = &po.User{
			AuId:     user.AuId,
			UserName: user.UserName,
		}
	}
	return list
}
