package test

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/peifengll/task-rebot/infra/component/mysql/ingorm"
	"github.com/peifengll/task-rebot/infra/persistence/dao"
	"github.com/peifengll/task-rebot/infra/persistence/po"
	"testing"
)

// 自动生成数据库
func TestAutoDb(t *testing.T) {
	db := ingorm.ProviderOnceGormDB()
	err := db.AutoMigrate(&po.User{}, &po.Service{}, &po.Task{}, &po.Subscribe{})
	if err != nil {
		log.Println(err)
		return
	}

}

func TestServiceDao(t *testing.T) {
	s := dao.ProviderServiceDaoInForm(ingorm.ProviderOnceGormDB())

	ctx := context.Background()
	service, _ := s.FindAllService(ctx)
	for _, v := range service {
		fmt.Printf("%+v\n", v)
	}
}

func TestCreateService(t *testing.T) {
	s := dao.ProviderServiceDaoInForm(ingorm.ProviderOnceGormDB())
	ctx := context.Background()
	for i := 0; i < 10; i++ {
		p := &po.Service{

			Name:       fmt.Sprintf("无敌%d", i),
			ServerTime: fmt.Sprintf("%02d:00", i),
			Status:     0,
		}
		err := s.CreateService(ctx, p)
		if err != nil {
			log.Fatalln(err)
		}
	}

}

// Users

func TestCreateUser(t *testing.T) {
	u := dao.ProviderUserDaoInGorm(ingorm.ProviderOnceGormDB())
	ctx := context.Background()
	for i := 0; i < 10; i++ {
		p := &po.User{
			AuId:     fmt.Sprintf("a%d", i),
			UserName: fmt.Sprintf("a%d", i),
		}
		err := u.CreateUser(ctx, p)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}
}
func TestFindNewId(t *testing.T) {
	u := dao.ProviderUserDaoInGorm(ingorm.ProviderOnceGormDB())
	ctx := context.Background()
	id, err := u.FindNewId(ctx)
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Printf("maxid is :%d\n", id)

}

// tasks

func TestCreateTask(t *testing.T) {
	taskDao := dao.ProviderTaskDao(ingorm.ProviderOnceGormDB())
	ctx := context.Background()
	for i := 0; i < 10; i++ {
		t := &po.Task{
			AuId:        "a0",
			Name:        fmt.Sprintf("测试%d", i),
			Description: fmt.Sprintf("描述%d", i),
			DueDate:     time.Now(),
			Status:      0,
		}
		err := taskDao.CreateTask(ctx, t)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}
}

// subscribes

func TestCreateSubscribe(t *testing.T) {
	s := dao.ProviderSubscribeDaoInGorm(ingorm.ProviderOnceGormDB())
	ctx := context.Background()
	for i := 0; i < 10; i++ {
		su := &po.Subscribe{
			AuId:      "a0",
			ServiceId: 0,
		}
		err := s.CreateSubscribe(ctx, su)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}
}

func TestGetSubscribeByServiceId(t *testing.T) {
	s := dao.ProviderSubscribeDaoInGorm(ingorm.ProviderOnceGormDB())
	ctx := context.Background()
	seid := 0
	auids, err := s.GetSubscribeByServiceId(ctx, seid)
	if err != nil {
		log.Fatalf("GetSubscribeByService id error %s", err)
		return
	}
	for _, v := range auids {
		fmt.Println(*v)
	}
}
