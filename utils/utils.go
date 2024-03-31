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

	if 6 <= hour && hour < 12 {
		return "Buenos días"
	}
	return ""
}

func getCurrentHour() int {
	dt := time.Now()
	return dt.Hour()
}
