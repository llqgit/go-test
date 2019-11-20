package main

import (
	"fmt"
	"github.com/robfig/cron"
	"os"
	"os/signal"
	"time"
)

// 定时任务测试

type TestCron struct {
}
type TestCron2 struct {
}

// 运行服务的方法
func (c *TestCron) Run() {
	t := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("run job1:", t)
}

func (c *TestCron2) Run() {
	t := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("run job2:", t)
}

func main() {
	// 创建cron tab定时任务
	c := cron.New()
	_ = c.AddJob("0 0 12 * * SUN", &TestCron{})
	_ = c.AddJob("2 * * * * SUN", &TestCron2{})
	//_ = c.AddJob(CronMonday, &services.MondayCron{})
	//_ = c.AddJob("0 * * * *", &services.WeekCron{}) // 测试
	c.Start()
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	for {
		select {
		case <-signals:
			return
		}
	}
}
