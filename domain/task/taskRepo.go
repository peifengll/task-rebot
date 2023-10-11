package task

import (
	"context"
)

type TaskRepo interface {
	SaveTask(ctx context.Context, task *Task) error
	UpdateTaskStatus(ctx context.Context, task *Task) error
	UpdateTaskDueDate(ctx context.Context, task *Task) error
	UpdateTaskDescription(ctx context.Context, task *Task) error
	UpdateTaskName(ctx context.Context, task *Task) error
	// 按状态查询今天到期的所有任务 ，-1就是不区分
	FindDueTasksByStatus(ctx context.Context, status int) ([]*Task, error)
	// 按状态查询今天到期的所有私人任务 ，-1就是不区分
	FindDueTaskPerson(ctx context.Context, auid string, status int) ([]*Task, error)
}
