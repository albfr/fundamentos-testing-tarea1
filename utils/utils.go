package utils

import "time"

func Hello(name string) string {

	var ct int = getCurrentHour()

	return getGreeting(ct)
}

func getGreeting(hour int) string {
	// [20, 6[
	// [6, 12[
	// [12, 20[

	return ""
}

func getCurrentHour() int {
	dt := time.Now()
	return dt.Hour()
}
