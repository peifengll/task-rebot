package monitor

import (
	"github.com/robfig/cron/v3"
)

type TimeHandler interface {
	//启动一个定时器，监视现在的时间， 在某个时间段就要做某些操作

}

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
