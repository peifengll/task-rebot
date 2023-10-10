package converter

import (
	"github.com/peifengll/task-rebot/domain/task"
	"github.com/peifengll/task-rebot/infra/persistence/po"
	"gorm.io/gorm"
)

type TaskConverter interface {
	ToEntityTask(task *po.Task) *task.Task
	ToEntityTasks(tasks []*po.Task) []*task.Task
	ToPoTask(task *task.Task) *po.Task
	ToPoTasks(tasks []*task.Task) []*po.Task
}

type TaskConverterImpl struct {
}

func NewTaskConverter() TaskConverter {
	return &TaskConverterImpl{}
}

var _ TaskConverter = &TaskConverterImpl{}

func (p *TaskConverterImpl) ToEntityTask(task *po.Task) *task.Task {
	return &task.Task{
		Id:          task.ID,
		AuId:        task.AuId,
		Name:        task.Name,
		Description: task.Description,
		DueDate:     task.DueDate,
		Status:      task.Status,
	}
}
func (p *TaskConverterImpl) ToEntityTasks(tasks []*po.Task) []*task.Task {
	enlistTasks := make([]*task.Task, len(tasks))
	for i, task := range tasks {
		enlistTasks[i] = &task.Task{
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
func (p *TaskConverterImpl) ToPoTask(task *task.Task) *po.Task {
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
func (p *TaskConverterImpl) ToPoTasks(tasks []*task.Task) []*po.Task {
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
