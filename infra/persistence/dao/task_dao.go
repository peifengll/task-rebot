package dao

import (
	"context"
	"github.com/peifengll/task-rebot/infra/persistence/po"
	"time"
)

type TaskDao interface {
	//	添加一个任务
	CreateTask(ctx context.Context, task *po.Task) error
	//	删除一个任务
	DeleteTask(ctx context.Context, id int) error
	//	修改任务的状态
	UpdateTaskStatus(ctx context.Context, id int, status int) error
	//	完成一个任务
	FinishTask(ctx context.Context, id int) error
	//修改任务的截至日期
	UpdateTaskDueDate(ctx context.Context, id int, duedate time.Time) error
	//	 修改任务的描述
	UpdateTaskDescription(ctx context.Context, id int, description string) error
	//	修改任务的名字
	UpdateTaskName(ctx context.Context, id int, name string) error
	//	 查询出所有到期任务
	FindAllDueTask(ctx context.Context) ([]*po.Task, error)
	//	查询出今天的个人任务
	FindAllDueTaskPersonal(ctx context.Context, auid string) ([]*po.Task, error)
	//		查询出所有未完成任务
	FindFinishDueTaskPersonal(ctx context.Context, auid string) ([]*po.Task, error)
}
