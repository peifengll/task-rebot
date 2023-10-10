package dao

import (
	"context"
	log "github.com/cihub/seelog"
	"github.com/peifengll/task-rebot/infra/persistence/po"
	"gorm.io/gorm"
)

type SubscribeDao interface {
	// 增加一条订阅
	CreateSubscribe(ctx context.Context, s *po.Subscribe) error
	//	删除一条订阅
	DeleteSubscribe(ctx context.Context, id int) error
	//	按找服务id查询一堆订阅,查询出用户id
	GetSubscribeByServiceId(ctx context.Context, seid int) ([]*string, error)
}

var _ SubscribeDao = &SubscribeDaoInGorm{}

type SubscribeDaoInGorm struct {
	db *gorm.DB
}

func NewSubscribeDaoInGorm(d *gorm.DB) *SubscribeDaoInGorm {
	return &SubscribeDaoInGorm{
		db: d,
	}
}

// 增加一条订阅
func (t *SubscribeDaoInGorm) CreateSubscribe(ctx context.Context, s *po.Subscribe) error {
	err := t.db.Model(&po.Subscribe{}).Create(s).Error
	if err != nil {
		log.Errorf("CreateSubscribe failed, err: %v", err)
		return err
	}
	return nil
}

// 删除一条订阅
func (t *SubscribeDaoInGorm) DeleteSubscribe(ctx context.Context, id int) error {
	err := t.db.Model(&po.Subscribe{}).Delete(&po.Subscribe{}, id).Error
	if err != nil {
		log.Errorf("DeleteSubscribe failed, err: %v", err)
		return err
	}
	return nil
}

// 按找服务id查询一堆订阅,查询出用户id
func (t *SubscribeDaoInGorm) GetSubscribeByServiceId(ctx context.Context, seid int) ([]*string, error) {
	subscribes := make([]*po.Subscribe, 0)
	err := t.db.Model(&po.Subscribe{}).Where("serviceid = ?", seid).Find(&subscribes).Error
	if err != nil {
		log.Errorf("GetSubscribeByServiceId failed, err: %v", err)
		return nil, err
	}
	var userids []*string
	for _, v := range subscribes {
		userids = append(userids, &v.AuId)
	}
	return userids, nil
}
