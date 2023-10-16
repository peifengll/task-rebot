package ticker

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

// 直接每过一个小时看看，要是当时快到 8点， 那就看 再开个任务定时器，去执行任务

func TimeCheck() {
	c := cron.New()
	_, err := c.AddFunc("@every 1h", func() {
		now := time.Now()
		fmt.Println(now.Format("15:04:05"))
	})
	if err != nil {
		return
	}
	c.Start()
	select {}
}
