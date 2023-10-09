package converter

import (
	"github.com/peifengll/task-rebot/domain/entity"
	"github.com/peifengll/task-rebot/infra/persistence/po"
)

type TaskConverter interface {
	ToEntityTask(task *po.Task) *entity.Task
	ToEntityTasks(tasks []*po.Task) []*entity.Task
	ToPoTask(task *entity.Task) *po.Task
	ToPoTasks(tasks []*entity.Task) []*po.Task
}
