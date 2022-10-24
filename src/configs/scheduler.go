package configs

import "os"

func SchedulerRun() string {
	return os.Getenv("SCHEDULER_RUN")
}
