package repositoryImpl

import (
	"context"
	log "github.com/cihub/seelog"
	"github.com/peifengll/task-rebot/domain/entity"
	"github.com/peifengll/task-rebot/domain/repository"
	"github.com/peifengll/task-rebot/infra/converter"
	"github.com/peifengll/task-rebot/infra/persistence/dao"
)

type UserRepoImpl struct {
	Dao     dao.UserDao
	Convert converter.UserConverter
}

var _ repository.UserRepo = &UserRepoImpl{}

func NewUserRepo(dao dao.UserDao, convert converter.UserConverter) repository.UserRepo {
	return &UserRepoImpl{
		Dao:     dao,
		Convert: convert,
	}
}

func (u *UserRepoImpl) SaveUser(ctx context.Context, user *entity.User) error {
	return u.Dao.CreateUser(ctx, u.Convert.ToPoUser(user))
}

// 获取一个新的id
func (u *UserRepoImpl) GetNewAuId(ctx context.Context) (int, error) {
	id, err := u.Dao.FindNewId(ctx)
	if err != nil {
		log.Errorf("Failed to get new uid: %v", err)
		return 0, err
	}
	return id, nil
}
