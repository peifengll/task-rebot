package converter

import (
	"github.com/peifengll/task-rebot/domain/entity"
	"github.com/peifengll/task-rebot/infra/persistence/po"
	"gorm.io/gorm"
)

type TaskConverter interface {
	ToEntityTask(task *po.Task) *entity.Task
	ToEntityTasks(tasks []*po.Task) []*entity.Task
	ToPoTask(task *entity.Task) *po.Task
	ToPoTasks(tasks []*entity.Task) []*po.Task
}

type TaskConverterImpl struct {
}

func NewTaskConverter() TaskConverter {
	return &TaskConverterImpl{}
}

var _ TaskConverter = &TaskConverterImpl{}

func (p *TaskConverterImpl) ToEntityTask(task *po.Task) *entity.Task {
	return &entity.Task{
		Id:          task.ID,
		AuId:        task.AuId,
		Name:        task.Name,
		Description: task.Description,
		DueDate:     task.DueDate,
		Status:      task.Status,
	}
}
func (p *TaskConverterImpl) ToEntityTasks(tasks []*po.Task) []*entity.Task {
	enlistTasks := make([]*entity.Task, len(tasks))
	for i, task := range tasks {
		enlistTasks[i] = &entity.Task{
			Id:          task.ID,
			AuId:        task.AuId,
			Name:        task.Name,
			Description: task.Description,
			DueDate:     task.DueDate,
			Status:      task.Status,
		}
	}
	return enlistTasks
}
func (p *TaskConverterImpl) ToPoTask(task *entity.Task) *po.Task {
	return &po.Task{
		Model: gorm.Model{
			ID: task.Id,
		},
		AuId:        task.AuId,
		Name:        task.Name,
		Description: task.Description,
		DueDate:     task.DueDate,
		Status:      task.Status,
	}
}
func (p *TaskConverterImpl) ToPoTasks(tasks []*entity.Task) []*po.Task {
	polistTasks := make([]*po.Task, len(tasks))
	for i, task := range tasks {
		polistTasks[i] = &po.Task{
			Model: gorm.Model{
				ID: task.Id,
			},
			AuId:        task.AuId,
			Name:        task.Name,
			Description: task.Description,
			DueDate:     task.DueDate,
			Status:      task.Status,
		}
	}
	return polistTasks
}
