package helper

import (
	"fmt"
	"time"
)

func GetDateTimeNow() string {
	currentTime := time.Now()

	date := fmt.Sprintf("%d-%d-%d %d:%d:%d\n",
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		currentTime.Hour(),
		currentTime.Hour(),
		currentTime.Second())

	return date
}
