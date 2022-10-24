package apps

import (
	"microservice/src/configs"
	"time"

	"github.com/go-co-op/gocron"
)

func SchedulerInit() *gocron.Scheduler {

	now, _ := time.LoadLocation(configs.GetTimezone())
	scheduler := gocron.NewScheduler(now)

	return scheduler
}
