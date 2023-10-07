package monitor

import (
	"github.com/robfig/cron/v3"
)

// 启动一个定时器，监视现在的时间
// 每小时，检测一次现在的时间，计算
func StartTimeMonitor() {
	c := cron.New()
	_, err := c.AddFunc("@every 1h", func() {
		//now:=time.Now()

	})
	if err != nil {
		panic(err)
		return
	}
	c.Start()
	select {}
}
