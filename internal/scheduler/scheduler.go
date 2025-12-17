package scheduler

import "time"

func Allowed(startHour, endHour int) bool {
	h := time.Now().Hour()
	return h >= startHour && h <= endHour
}
