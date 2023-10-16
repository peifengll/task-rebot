package test

import (
	"fmt"
	"github.com/peifengll/task-rebot/config"
	"github.com/peifengll/task-rebot/domain/entity"
	"github.com/peifengll/task-rebot/domain/repo"
	"github.com/peifengll/task-rebot/infra/persistence/task"
	"testing"
	"time"
)

var taskDao repo.TaskRepo = &task.TaskRepo{}

func TestAddTask(t *testing.T) {
	config.InitDb()
	taskList := make([]*entity.Task, 0)
	for i := 0; i < 5; i++ {
		taskList = append(taskList, &entity.Task{
			AuId:        "a_zhangsan",
			Name:        "占领日本" + string(rune(i)),
			Description: "干就完了",
			DueDate:     time.Now(),
			Status:      0,
		})
	}
	for i := 0; i < len(taskList); i++ {
		taskDao.AddTask(taskList[i])

	}

}

func TestDueTimeQB(t *testing.T) {
	config.InitDb()
	tasks, err := taskDao.GetDueTasks()
	if err != nil {
		return
	}
	for i := 0; i < len(tasks); i++ {
		fmt.Println(tasks[i].ID, tasks[i].DueDate, tasks[i].Name,
			tasks[i].Description, tasks[i].Status)
	}
}

func TestDueTimeQByID(t *testing.T) {
	config.InitDb()
	tasks, err := taskDao.GetDueTasksByAuId("a_zhangsan")
	if err != nil {
		return
	}
	for i := 0; i < len(tasks); i++ {
		fmt.Println(tasks[i].ID, tasks[i].DueDate, tasks[i].Name,
			tasks[i].Description, tasks[i].Status)
	}
}
