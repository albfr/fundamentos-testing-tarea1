package utils

import (
	"time"
)

func IsPalindome(s string) bool {
	return false
}

func ReverseString(name string) string {
	r := []rune(name)
	n := len(name)
	for i := 0; i < len(name)/2; i++ {
		r[i], r[n-1-i] = r[n-1-i], r[i]
	}
	return string(r)
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
