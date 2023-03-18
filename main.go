package main

import (
	"todo_list/conf"
	"todo_list/routs"
)

func main() {
	conf.Init()
	r := routs.CreateRoutes()
	r.Run(conf.AppPort)
}
