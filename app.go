package main

import (
	"webgin/views"
	"webgin/model"
)

func main() {
	defer model.MasterDB.Close()
	views.Start()
}