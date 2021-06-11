package tasks

import (
	"fmt"
	"tdapi/dataservice"

	"github.com/jakecoffman/cron"
)

var task *cron.Cron

type TestJob struct {
}

func (t TestJob) Run() {
	fmt.Println("testJob1...")
}

type Test2Job struct {
}

func (t Test2Job) Run() {
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

func InsertTask() {
}

func CloseTask() {
	//关闭着计划任务, 但是不能关闭已经在执行中的任务.
	defer task.Stop()
}
