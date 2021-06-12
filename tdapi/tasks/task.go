package tasks

import (
	"fmt"
	"tdapi/dataservice"
	"tdapi/log"
	"tdapi/model"
	"time"

	"github.com/robfig/cron/v3"
)

var task *cron.Cron

const (
	IsCycle = 1
)

type Job struct {
	Tid        int
	Account    int
	Groupid    string
	NeedCounts int
	Countsed   int //已经执行次数
	Cron       int
	Cycle      int
	Text       string
	Shut       chan int
	Entry      cron.EntryID
}

func (t Job) Run() {

	if t.NeedCounts == 0 {
		task.Remove(t.Entry)
		log.Info("清理定时器", t.Account, t.Tid)
	}
	if t.Cycle != IsCycle {
		task.Remove(t.Entry)
		log.Info("清理定时器", t.Account, t.Tid)
	}
	fmt.Println(t.Tid, t.Account, t.Text, time.Now())
	t.NeedCounts--
	t.Countsed++
}

func InitTasks() {
	task = cron.New()

	t, err := dataservice.LoadTaks()
	fmt.Println(t, err)
	InsertTask(t)
	//启动计划任务
	task.Start()

	select {}

}

func InsertTask(t []model.Task) {
	for _, value := range t {
		var spec string

		if value.Cycle == IsCycle {
			spec = fmt.Sprintf("* */%d * * *", value.Cron)
		}

		j := Job{
			Tid:        value.Tid,
			Account:    value.Account,
			Groupid:    value.Groupid,
			NeedCounts: value.Counts,
			Cycle:      value.Cycle, //需要循环
			Text:       value.Text,
			Shut:       make(chan int, 1),
		}
		task.AddJob(spec, &j)

		fmt.Println(value, j, spec)
	}
}

func CloseTask() {
	//关闭着计划任务, 但是不能关闭已经在执行中的任务.
	defer task.Stop()
}
