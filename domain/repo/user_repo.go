package repo

import "github.com/peifengll/task-rebot/domain/entity"

type UserRepo interface {
	// Get 查这个人的相关信息
	Get(string) (*entity.User, error)
	// Save 更新还是保存
	Save(*entity.User) (bool, error)
}
