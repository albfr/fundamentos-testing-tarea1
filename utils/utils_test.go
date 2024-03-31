package utils

import "testing"

func TestGetGreeting(t *testing.T) {
	want := "Buenos días"
	hour := 8
	if got := getGreeting(hour); got != want {
		t.Errorf("getGreeting(%d) = \"%s\", want = %s", hour, got, want)
	}
}
