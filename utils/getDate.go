package utils

import (
	"time"
)

func GetDate(num int) []string {
	var result []string
	currentTime := time.Now()
	currentTime.AddDate(0, 0, -1)
	for i := 0; i < num; i++ {
		result = append(result, currentTime.AddDate(0, 0, -i).Format("2006-01-02"))
	}
	return result
}
