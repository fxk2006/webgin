package main

import (
	_ "webgin/views"
	_ "webgin/model"
	"webgin/views"
	"webgin/model"
)

func main() {
	defer model.DB.Close()
	views.Engine.Run()
}