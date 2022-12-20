package helper

import (
	"fmt"
	"time"
)

func GetDateTimeNowString() string {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	currentTime := time.Now().In(loc)

	date := fmt.Sprintf("%d-%d-%d %d:%d:%d\n",
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		currentTime.Hour(),
		currentTime.Hour(),
		currentTime.Second())

	return date
}

func GetDateTimeNowDate() time.Time {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	currentTime := time.Now().In(loc)
	return currentTime
}
