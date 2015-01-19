package handlers

import (
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func UptimeHandler(msg string) string {
	elapsed := time.Since(start)
	return elapsed.String()
}
