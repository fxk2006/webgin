package tasks

import (
	"github.com/satori/go.uuid"
	"webgin/global"
	"os/exec"
	"io"
	"bufio"
)

const (
	Running = iota
	Success
	Failure
)

type Progress struct {
	TaskStatus byte
	CurrentPos int
	Message    string
}


type Tasks map[uuid.UUID]*string

var DBTasks = make(Tasks)

func ExecCommand(cmd string) (result io.ReadCloser,pid int, err error) {
	command := exec.Command("/bin/bash", "-c", cmd)
	result, err = command.StdoutPipe()
	if err := command.Start();err != nil {
		global.GLog.Error(err)
	}
	pid = command.Process.Pid
	go func() {
		err :=command.Wait()
		if err != nil{
			global.GLog.Debug("子任务进程",pid,"异常退出",err)
		}else {
			global.GLog.Error("子任务进程",pid,"正常退出")
		}
	}()
	return
}

func LongTask(id uuid.UUID) {
	global.GLog.Debug(id, "already start ...")
	//defer delete(Ttasks,id)
	var out string
	DBTasks[id] = &out
	cmd := `./longtask.sh`
	stdout,pid,_ := ExecCommand(cmd)
	global.GLog.Debug(id, "子任务进程 already running  pid",pid)
	reader := bufio.NewReader(stdout)
	for {
		line,_,err1 := reader.ReadLine()
		out += string(line)
		out +="\n"
		if err1 != nil || err1 == io.EOF{
			break
		}
	}
	global.GLog.Debug(id, pid,"子任务进程 already finish ...")
}
