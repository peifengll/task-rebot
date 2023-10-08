package test

import (
	"fmt"
	"github.com/peifengll/task-rebot/infra/component/config"
	"testing"
	"time"
)

func TestGetConfig(t *testing.T) {
	time.Sleep(time.Second * 3)
	c, err := config.GetConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", c)
}
