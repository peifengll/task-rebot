package dao

import (
	"context"
	log "github.com/cihub/seelog"
	"github.com/peifengll/task-rebot/infra/persistence/po"
	"gorm.io/gorm"
)

// 大体就是服务的增删查改功能
type ServiceDao interface {
	//按照id查询service
	FindServiceById(ctx context.Context, id int) (*po.Service, error)
	//查询所有服务
	FindAllService(ctx context.Context) ([]*po.Service, error)
	//	按状态查询服务
	FindAllServiceStatus(ctx context.Context, status int) ([]*po.Service, error)
	//	按照该服务启动的时间查询服务
	FindServicesByServerTime(ctx context.Context, servertime string) ([]*po.Service, error)
	//	添加服务记录
	CreateService(ctx context.Context, service *po.Service) error
	//修改服务是否启用
	UpdateServiceStatus(ctx context.Context, id int, status int) error
	//	删除一个服务
	DeleteService(ctx context.Context, id int) error
}

func NewServiceDaoInForm(d *gorm.DB) *ServiceDaoInForm {
	return &ServiceDaoInForm{
		db: d,
	}
}

type ServiceDaoInForm struct {
	db *gorm.DB
}

var _ ServiceDao = &ServiceDaoInForm{}

func (r *ServiceDaoInForm) FindServiceById(ctx context.Context, id int) (*po.Service, error) {
	p := &po.Service{}
	err := r.db.Model(&po.Service{}).Where("id = ?", id).First(p).Error
	if err != nil {
		log.Errorf("error querying service FindServiceById:%s", err)
		return nil, err
	}
	return p, nil
}

func (r *ServiceDaoInForm) FindAllService(ctx context.Context) ([]*po.Service, error) {
	plist := make([]*po.Service, 0)
	err := r.db.Model(&po.Service{}).Find(&plist).Error
	if err != nil {
		log.Errorf("error querying service FindAllService:%s", err)
		return nil, err
	}
	return plist, nil
}

func (r *ServiceDaoInForm) FindAllServiceStatus(ctx context.Context, status int) ([]*po.Service, error) {
	plist := make([]*po.Service, 0)
	var err = r.db.Model(&po.Service{}).Where("status = ?", status).Find(&plist).Error
	if err != nil {
		log.Errorf("error querying service FindAllServiceStatus:%s", err)
		return nil, err
	}
	return plist, nil
}

func (r *ServiceDaoInForm) FindServicesByServerTime(ctx context.Context, servertime string) ([]*po.Service, error) {
	plist := make([]*po.Service, 0)
	err := r.db.Model(&po.Service{}).Where("servertime = ?", servertime).Find(&plist).Error
	if err != nil {
		log.Errorf("error querying service FindServicesByServerTime:%s", err)
		return nil, err
	}
	return plist, nil
}
func (r *ServiceDaoInForm) CreateService(ctx context.Context, service *po.Service) error {
	err := r.db.Create(service).Error
	if err != nil {
		log.Errorf("error create service CreateService:%s", err)
		return err
	}
	return nil
}
func (r *ServiceDaoInForm) UpdateServiceStatus(ctx context.Context, id int, status int) error {
	err := r.db.Model(&po.Service{}).Where("id = ?", id).Update("status", status).Error
	if err != nil {
		log.Errorf("error updating service status:%s", err)
		return err
	}
	return nil
}
func (r *ServiceDaoInForm) DeleteService(ctx context.Context, id int) error {
	err := r.db.Model(&po.Service{}).Delete(&po.Service{}, id).Error
	if err != nil {
		log.Errorf("error deleting service:%s", err)
		return err
	}
	return nil
}
