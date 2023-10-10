package repositoryImpl

import (
	"context"
	log "github.com/cihub/seelog"
	"github.com/peifengll/task-rebot/domain/entity"
	"github.com/peifengll/task-rebot/domain/repository"
	"github.com/peifengll/task-rebot/infra/converter"
	"github.com/peifengll/task-rebot/infra/persistence/dao"
)

type SubscribeRepoImpl struct {
	Dao     dao.SubscribeDao
	Convert converter.SubscribeConverter
}

func NewSubscribeRepo(dao dao.SubscribeDao, convert converter.SubscribeConverter) repository.SubscribeRepo {
	return &SubscribeRepoImpl{
		Dao:     dao,
		Convert: convert,
	}
}

var _ repository.SubscribeRepo = &SubscribeRepoImpl{}

// 按照服务id 查询信息，
func (p *SubscribeRepoImpl) FindUserInfo(ctx context.Context, serviceId int) ([]*string, error) {

	auids, err := p.Dao.GetSubscribeByServiceId(ctx, serviceId)
	if err != nil {
		log.Errorf("get subscribe by service id %d error: %v", serviceId, err)
		return nil, err
	}
	return auids, err
}

// 添加一个
func (p *SubscribeRepoImpl) AddOneRecord(ctx context.Context, subscribe *entity.Subscribe) error {
	po := p.Convert.TpPoSubscribe(subscribe)
	err := p.Dao.CreateSubscribe(ctx, po)
	if err != nil {
		log.Errorf("create subscribe error: %v", err)
		return err
	}
	return nil
}
