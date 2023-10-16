package task

import (
	"github.com/peifengll/task-rebot/config"
	"github.com/peifengll/task-rebot/domain/entity"
)

type TaskRepo struct {
}

// todo 具体逻辑 返回到期的所有任务
func (r *TaskRepo) GetDueTasks() ([]*entity.Task, error) {
	tl := make([]*entity.Task, 0)
	err := config.GetDb().Model(&entity.Task{}).
		Where("date(due_date) = date(now())").Find(&tl).Error
	if err != nil {
		return nil, err
	}
	return tl, err
}

// 获取某个用户今日到期的任务
func (r *TaskRepo) GetDueTasksByAuId(auid string) ([]*entity.Task, error) {
	tl := make([]*entity.Task, 0)
	err := config.GetDb().Model(&entity.Task{}).
		Where("au_id = ? and  date(due_date) = date(now())", auid).Find(&tl).Error
	if err != nil {
		return nil, err
	}
	return tl, err
}

func (r *TaskRepo) AddTask(z *entity.Task) error {
	err := config.GetDb().Model(&entity.Task{}).Save(z).Error
	return err
}

func (r *TaskRepo) SolveTaskById(id uint, op int) error {

	return nil
}
