package dao

import (
	"context"
	log "github.com/cihub/seelog"
	"github.com/peifengll/task-rebot/infra/persistence/po"
	"gorm.io/gorm"
)

type UserDao interface {
	// 返回一个id
	CreateUser(ctx context.Context, user *po.User) error

	//	找到新的编号
	FindNewId(ctx context.Context) (int, error)
}

type UserDaoInGorm struct {
	db *gorm.DB
}

func ProviderUserDaoInGorm(d *gorm.DB) *UserDaoInGorm {
	return &UserDaoInGorm{
		db: d,
	}
}

func (u *UserDaoInGorm) CreateUser(ctx context.Context, user *po.User) error {
	err := u.db.Model(&po.User{}).Create(user).Error
	if err != nil {
		log.Errorf("Error creating user:%s", err)
		return err
	}
	return nil
}

func (u *UserDaoInGorm) FindNewId(ctx context.Context) (int, error) {
	row := u.db.Raw("SELECT MAX(id) FROM users").Row()
	var id int
	err := row.Scan(&id)
	if err != nil {
		log.Errorf("Error: %s", err)
		return 0, err
	}
	return id, nil

}
