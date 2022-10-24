package scheduler

import (
	"microservice/src/apps"
	"microservice/src/configs"
	"microservice/src/scheduler/tasks"
)

// reference : https://github.com/go-co-op/gocron

func Register() {
	if configs.SchedulerRun() == "true" {
		scheduler := apps.SchedulerInit()

		//-> Register here...
		scheduler.Every(10).Seconds().WaitForSchedule().Do(tasks.Example)

		scheduler.StartAsync()
	}
}
