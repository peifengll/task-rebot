package test

import (
	"github.com/peifengll/task-rebot/infra/component/monitor"
	"testing"
)

func TestName(t *testing.T) {
	monitor.StartTimeMonitor()
}
