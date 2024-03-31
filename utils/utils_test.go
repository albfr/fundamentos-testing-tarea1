package utils

import "testing"

type addTest struct {
	arg1     int
	expected string
}

var addTests = []addTest{
	{8, "Buenos días"},
	{13, "Buenas tardes"},
	{22, "Buenas noches"},
}

func TestGetGreeting(t *testing.T) {

	for _, test := range addTests {
		if got := getGreeting(test.arg1); got != test.expected {
			t.Errorf("getGreeting(%d) = \"%s\", want = \"%s\"", test.arg1, got, test.expected)
		}
	}

}
