package main

import (
	_ "webgin/model"
	"webgin/views"
	"webgin/model"
)

func main() {
	defer model.MasterDB.Close()
	views.Start()
}