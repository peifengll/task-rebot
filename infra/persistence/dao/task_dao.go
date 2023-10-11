package dao

import (
	"context"
	log "github.com/cihub/seelog"
	"github.com/peifengll/task-rebot/infra/persistence/po"
	"gorm.io/gorm"
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

	//	 按照状态查询出今天的个人任务
	FindDueTaskByStatusPersonal(ctx context.Context, auid string, status int) ([]*po.Task, error)
	//	 按照状态查询今天的所有任务
	FindDueTaskByStatus(ctx context.Context, status int) ([]*po.Task, error)
	//	 按照状态查询所有个人任务
	FindTasksByStatusPersonal(ctx context.Context, auid string, status int) ([]*po.Task, error)
}

type TaskDaoInGorm struct {
	db *gorm.DB
}

func NewTaskDao(db *gorm.DB) TaskDao {
	return &TaskDaoInGorm{
		db: db,
	}
}

// 添加一个任务
func (t *TaskDaoInGorm) CreateTask(ctx context.Context, task *po.Task) error {
	err := t.db.Model(&po.Task{}).Create(&task).Error
	if err != nil {
		log.Errorf("Task creation error %v", err)
		return err
	}
	return nil
}

// 删除一个任务
func (t *TaskDaoInGorm) DeleteTask(ctx context.Context, id int) error {
	err := t.db.Model(&po.Task{}).Delete(&po.Task{}, id).Error
	if err != nil {
		log.Errorf("Error deleting task  %+v %+v", id, err)
		return err
	}
	return nil
}

// 修改任务的状态
func (t *TaskDaoInGorm) UpdateTaskStatus(ctx context.Context, id int, status int) error {
	err := t.db.Model(&po.Task{}).Where("id = ?", id).Update("status", status).Error
	if err != nil {
		log.Errorf("Error updating task  %+v %+v", id, err)
		return err
	}
	return nil
}

// 完成一个任务
func (t *TaskDaoInGorm) FinishTask(ctx context.Context, id int) error {
	//	 完成一个任务就是把该任务的状态设置为1
	return t.UpdateTaskStatus(ctx, id, 1)
}

// 修改任务的截至日期
func (t *TaskDaoInGorm) UpdateTaskDueDate(ctx context.Context, id int, duedate time.Time) error {
	err := t.db.Model(&po.Task{}).Where("id = ?", id).Update("duedate", duedate).Error
	if err != nil {
		log.Errorf("Error updating task due date %v: %v", id, err)
		return err
	}
	return nil
}

// 修改任务的描述
func (t *TaskDaoInGorm) UpdateTaskDescription(ctx context.Context, id int, description string) error {
	err := t.db.Model(&po.Task{}).Where("id = ?", id).Update("description", description).Error
	if err != nil {
		log.Errorf("Error updating task description %v: %v", id, err)
		return err
	}
	return nil
}

// 修改任务的名字
func (t *TaskDaoInGorm) UpdateTaskName(ctx context.Context, id int, name string) error {
	err := t.db.Model(&po.Task{}).Where("id = ?", id).Update("name", name).Error
	if err != nil {
		log.Errorf("Error updating task name: %v", err)
		return err
	}
	return nil
}

// 查询出所有到期任务
func (t *TaskDaoInGorm) FindAllDueTask(ctx context.Context) ([]*po.Task, error) {
	tasks := make([]*po.Task, 0)
	//select * from tasks where to_days(duetime) = to_days(now());
	err := t.db.Model(&po.Task{}).Where("to_days(duedate) = to_days(now())").Find(&tasks).Error
	if err != nil {
		log.Errorf("Error while querying AllDueTasks: %v", err)
		return nil, err
	}
	return tasks, nil
}

// 查询出今天的个人任务
func (t *TaskDaoInGorm) FindAllDueTaskPersonal(ctx context.Context, auid string) ([]*po.Task, error) {
	tasks := make([]*po.Task, 0)
	err := t.db.Model(&po.Task{}).Where("to_days(duedate) = to_days(now()) and auid = ?", auid).Find(&tasks).Error
	if err != nil {
		log.Errorf("Error while querying Personal AllDueTasks: %v", err)
		return nil, err
	}
	return tasks, nil

}

// 查询出个人所有未完成任务
func (t *TaskDaoInGorm) FindFinishDueTaskPersonal(ctx context.Context, auid string) ([]*po.Task, error) {
	tasks := make([]*po.Task, 0)
	err := t.db.Model(&po.Task{}).Where("status == ? and auid = ?", 0, auid).Find(&tasks).Error
	if err != nil {
		log.Errorf("Error while querying Personal 0Tasks: %v", err)
		return nil, err
	}
	return tasks, nil
}

// 按照状态查询出今天的个人任务
func (t *TaskDaoInGorm) FindDueTaskByStatusPersonal(ctx context.Context, auid string, status int) (aim []*po.Task, err error) {
	err = t.db.Model(&po.Task{}).Where("to_days(duedate) = to_days(now()) and auid = ? and status = ?", auid, status).Find(&aim).Error
	if err != nil {
		log.Errorf("Failed to find task in db: %v", err)
		return nil, err
	}
	return aim, nil
}

// 按照状态查询今天的所有任务
func (t *TaskDaoInGorm) FindDueTaskByStatus(ctx context.Context, status int) (aim []*po.Task, err error) {
	err = t.db.Model(&po.Task{}).Where("status = ?", status).Find(&aim).Error
	if err != nil {
		log.Errorf(" could not find task in database  for status %v", err)
		return nil, err
	}
	return aim, nil
}

// 按照状态查询所有个人任务
func (t *TaskDaoInGorm) FindTasksByStatusPersonal(ctx context.Context, auid string, status int) (aim []*po.Task, err error) {
	err = t.db.Model(&po.Task{}).Where("auid=?", auid).Find(&aim).Error
	if err != nil {
		log.Errorf("Error while querying Personal Tasks: %v", err)
		return nil, err
	}
	return aim, nil
}
