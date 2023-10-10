package repositoryImpl

import (
	"context"
	log "github.com/cihub/seelog"
	"github.com/peifengll/task-rebot/domain/task"
	"github.com/peifengll/task-rebot/infra/converter"
	"github.com/peifengll/task-rebot/infra/persistence/dao"
)

type TaskRepoImpl struct {
	Dao     dao.TaskDao
	Convert converter.TaskConverter
}

var _ task.TaskRepo = &TaskRepoImpl{}

func NewTaskRepo(dao dao.TaskDao, convert converter.TaskConverter) task.TaskRepo {
	return &TaskRepoImpl{
		Dao:     dao,
		Convert: convert,
	}
}

func (t *TaskRepoImpl) SaveTask(ctx context.Context, task *task.Task) error {
	err := t.Dao.CreateTask(ctx, t.Convert.ToPoTask(task))
	if err != nil {
		log.Errorf("Error saving task %v: %v", task, err)
		return err
	}
	return nil
}

func (t *TaskRepoImpl) UpdateTaskStatus(ctx context.Context, task *task.Task) error {
	err := t.Dao.UpdateTaskStatus(ctx, int(task.Id), task.Status)
	if err != nil {
		log.Errorf("Error updating task status: %v", err)
		return err
	}
	return nil
}

func (t *TaskRepoImpl) UpdateTaskDueDate(ctx context.Context, task *task.Task) error {
	err := t.Dao.UpdateTaskDueDate(ctx, int(task.Id), task.DueDate)
	if err != nil {
		log.Errorf("Error updating task due date: %v", err)
		return err
	}
	return nil
}

func (t *TaskRepoImpl) UpdateTaskDescription(ctx context.Context, task *task.Task) error {
	err := t.Dao.UpdateTaskDescription(ctx, int(task.Id), task.Description)
	if err != nil {
		log.Errorf("Error updating task description: %v", err)
		return err
	}
	return nil
}

func (t *TaskRepoImpl) UpdateTaskName(ctx context.Context, task *task.Task) error {
	err := t.Dao.UpdateTaskName(ctx, int(task.Id), task.Name)
	if err != nil {
		log.Errorf("Error updating task name: %v", err)
		return err
	}
	return nil
}

// 按状态查询今天到期的所有任务 ，-1就是不区分
func (t *TaskRepoImpl) FindDueTasksByStatus(ctx context.Context, status int) ([]*task.Task, error) {
	poList, err := t.Dao.FindDueTaskByStatus(ctx, status)
	if err != nil {
		log.Errorf("Failed to find due tasks by status %d: %v", status, err)
		return nil, err
	}

	return t.Convert.ToEntityTasks(poList), nil
}

// 按状态查询今天到期的所有私人任务 ，-1就是不区分
func (t *TaskRepoImpl) FindDueTaskPerson(ctx context.Context, auid string, status int) ([]*task.Task, error) {
	polist, err := t.Dao.FindDueTaskByStatusPersonal(ctx, auid, status)
	if err != nil {
		log.Errorf("Failed to find due tasks by status %d: %v", status, err)
		return nil, err
	}
	return t.Convert.ToEntityTasks(polist), nil
}
