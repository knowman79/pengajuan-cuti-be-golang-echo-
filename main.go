package main

import (
	"example/routes"
	"example/schedule"
	"log"

	"github.com/prprprus/scheduler"
)

func main() {
	// Scheduler Checking and Execution
	s, err := scheduler.NewScheduler(0)
	if err != nil {
		log.Println(err.Error())
	}
	s.Every().Second(0).Minute(0).Day(1).Month(1).Do(schedule.CrossYearLeaveScheduler)
	s.Every().Second(0).Minute(0).Hour(0).Day(1).Month(7).Do(schedule.ExpiredLeaveScheduler)

	// Running Enpoint
	routes.Endpoint()
}
