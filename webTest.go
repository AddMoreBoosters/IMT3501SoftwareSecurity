package main 

import (
	"MVC"
)

func main() {
	model := MVC.Model{}
	view := MVC.View{}
	controller := MVC.Controller{}
	controller.SetModel(&model)
	controller.SetView(&view)

	controller.HandleWebTraffic()
}