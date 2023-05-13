package main

import (
	"todo_list/initialize"
	"todo_list/router"
)

func main() {
	initialize.Init()
	r := router.NewRouter()
	r.Run()
}
