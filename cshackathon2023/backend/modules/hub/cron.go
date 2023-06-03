package ihub

import (
	"github.com/go-co-op/gocron"
	"time"
)

func Cron() {

	s := gocron.NewScheduler(time.UTC)
	_, _ = s.Every(10 * time.Second).Do(CronBackdrop)
	s.StartAsync()
}
