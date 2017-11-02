package main

import "github.com/astaxie/beego"

type Main struct {
	beego.Controller
}

func (this Main)Get()  {
	this.Ctx.WriteString("hello world")
}
func main(){
	beego.Router("/",&Main{})
    beego.Run()
}