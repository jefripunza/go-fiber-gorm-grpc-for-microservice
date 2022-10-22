package scheduler

import (
	"main-service/src/scheduler/tasks"
	"time"

	"github.com/go-co-op/gocron"
)

// reference : https://github.com/go-co-op/gocron

func Register() {

	now, _ := time.LoadLocation("Asia/Jakarta")
	s := gocron.NewScheduler(now)

	//-> Register here...
	s.Every(10).Seconds().Do(tasks.Example)

	s.StartAsync()

}
