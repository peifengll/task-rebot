package repo

import "github.com/peifengll/task-rebot/domain/entity"

type TaskRepo interface {
	//   所有今天到期的任务
	GetDueTasks() ([]*entity.Task, error)
	// 某个用户今天到期的任务
	GetDueTasksByAuId(auid string) ([]*entity.Task, error)
	// 添加一个任务
	AddTask(z *entity.Task) error
	// 完成一个
	SolveTaskById(uint, int) error
}
