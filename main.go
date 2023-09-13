package main

import "api-trivia/service"

func main() {
	m := service.NewMonitorService()
	m.TakeDate()
	m.Observe(11)
}
