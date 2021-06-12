package tasks

import (
	"fmt"
	"tdapi/dataservice"

	"github.com/jakecoffman/cron"
)

var task *cron.Cron

const (
	IsCycle = 1
)

type Job struct {
	Tid     int
	Account int
	Groupid string
	Counts  string
	Cron    int
	Cycle   int
	Text    string
	Shut    chan int
}

func (t Job) Run() {
	fmt.Println("testJob2...")
}

func InitTasks() {
	task = cron.New()

	t, err := dataservice.LoadTaks()
	fmt.Println(t, err)

	//启动计划任务
	task.Start()

	select {}

}
func LoadTasks() {

}

func InsertTask(t []Job) {
	for _, value := range t {
		var spec string

		if value.Cycle == IsCycle {
			spec = fmt.Sprintf("*/%d * * * *", value.Cron)
		}

		j := Job{
			Tid:     value.Tid,
			Account: value.Account,
			Groupid: value.Groupid,
			Counts:  value.Counts,

			Cycle: value.Cycle,
			Text:  value.Text,
			Shut:  make(chan int, 1),
		}

		fmt.Println(value, j, spec)
	}
}

func CloseTask() {
	//关闭着计划任务, 但是不能关闭已经在执行中的任务.
	defer task.Stop()
}
