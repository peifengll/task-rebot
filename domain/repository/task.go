package repository

import (
	"context"
	"github.com/peifengll/task-rebot/domain/entity"
)

type TaskRepo interface {
	SaveTask(ctx context.Context, task *entity.Task) error

	UpdateTask(ctx context.Context, task *entity.Task) error
	// 按状态查询今天到期的所有任务 ，-1就是不区分
	FindDueTasksByStatus(ctx context.Context, status int) error
	// 按状态查询今天到期的所有私人任务 ，-1就是不区分
	FindDueTaskPerson(ctx context.Context, auid string, status int) (*entity.Task, error)
}
