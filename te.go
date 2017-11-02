package main

import (
    "fmt"
    "os/exec"
)

func main() {
    shellPath := "./1.sh"
    command := exec.Command(shellPath) //初始化Cmd
    err := command.Start()//运行脚本
    if nil != err {
        fmt.Println(err)
    }
    fmt.Println("Process PID:", command.Process.Pid)
    err = command.Wait() //等待执行完成
    if nil != err {
        fmt.Println(err)
    }
    fmt.Println("ProcessState PID:", command.ProcessState.Pid())
}

