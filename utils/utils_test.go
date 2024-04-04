package utils

import (
	"testing"
)

type addReverseStringTest struct {
	s        string
	expected string
}

var addReverseStringTests = []addReverseStringTest{
	{"hola", "aloh"},
}

type addGreetingTest struct {
	arg1     int
	expected string
}

var addGreetingTests = []addGreetingTest{
	{8, "Buenos días"},
	{13, "Buenas tardes"},
	{22, "Buenas noches"},
}

func TestGetGreeting(t *testing.T) {

	for _, test := range addGreetingTests {
		got := getGreeting(test.arg1)
		if got != test.expected {
			t.Errorf("getGreeting(%d) = \"%q\", want = \"%s\"", test.arg1, got, test.expected)
		}
	}

}
