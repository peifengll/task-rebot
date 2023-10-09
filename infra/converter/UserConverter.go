package converter

import (
	"github.com/peifengll/task-rebot/domain/entity"
	"github.com/peifengll/task-rebot/infra/persistence/po"
)

type UserConverter interface {
	ToEntityUser(user *po.User) *entity.User
	ToEntityUsers(user []*po.User) []*entity.User
	ToPoUser(user *entity.User) *po.User
	ToPoUsers(user []*entity.User) []*po.User
}
