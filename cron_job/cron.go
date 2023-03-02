package cron_job

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func hello(name string) {
	message := fmt.Sprintf("Hi, %v", name)
	fmt.Println(message)
}

func RunCronJobs() {
	s := cron.New()

	s.AddFunc("@every 1s", func() {
		hello("John Doe")
	})

	s.Start()
}
