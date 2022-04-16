package usecase

import (
	"log"
	"time"
)

func checkTime(t time.Time) bool {
	nowduration := time.Since(t)
	duration, err := time.ParseDuration("1m")
	if err != nil {
		log.Fatalln(err)
	}
	return nowduration >= duration
}
