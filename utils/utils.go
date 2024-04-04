package utils

import (
	"time"
)

func ReverseString(name string) string {
}

func Hello(name string) string {

	var ct int = getCurrentHour()

	return "¡" + getGreeting(ct) + " " + name + "!"
}

func getGreeting(hour int) string {
	// [20, 6[
	// [6, 12[
	// [12, 20[

	if 6 <= hour && hour < 12 {
		return "Buenos días"
	}

	if 12 <= hour && hour < 20 {
		return "Buenas tardes"
	}

	return "Buenas noches"
}

func getCurrentHour() int {
	dt := time.Now()
	return dt.Hour()
}
