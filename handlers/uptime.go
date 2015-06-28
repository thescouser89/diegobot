package handlers

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func UptimeHandler(msg string) string {
	elapsed := time.Since(start)
	if elapsed.Hours() > 48 {
		return fmt.Sprintf("%v days", elapsed.Hours()/24)
	} else {
		return elapsed.String()
	}
}
