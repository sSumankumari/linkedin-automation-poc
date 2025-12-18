package scheduler

import (
	"fmt"
	"time"
)

func Allowed(startHour, endHour int) bool {
	current := time.Now().Hour()
	fmt.Printf("[DEMO] Scheduler check: current=%d, allowed=%d-%d\n",
		current, startHour, endHour)

	return current >= startHour && current <= endHour
}
