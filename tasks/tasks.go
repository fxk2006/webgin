package tasks

import (
	"fmt"
	"time"
	"github.com/satori/go.uuid"
	"webgin/global"
)

const (
	Running = iota
	Success
	Failure
)

type Progress struct {
	TaskStatus byte
	CurrentPos int
}

type Tasks map[uuid.UUID] *Progress

var DBTasks = make(Tasks)

func LongTask(id uuid.UUID) {
	global.GLog.Debug(id,"already start ...")
	//defer delete(Ttasks,id)
	p := new(Progress)
	p.TaskStatus = Running
	DBTasks[id] = p
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
		p.CurrentPos = i
		time.Sleep(1 * time.Second)
	}
	p.TaskStatus = Success
	global.GLog.Debug(id,"%s already finish ...")
}
