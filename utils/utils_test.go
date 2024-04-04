package utils

import (
	"strconv"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	s := "a"
	got := IsPalindome(s)
	want := true
	str_got := strconv.FormatBool(got)
	str_want := strconv.FormatBool(want)

	if got != want {
		t.Errorf("IsPalindrome(\"%s\") = %s, want = %s", s, str_got, str_want)
	}

}

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

func TestReverseString(t *testing.T) {

	for _, test := range addReverseStringTests {
		got := ReverseString(test.s)
		want := test.expected
		if got != want {
			t.Errorf("ReverseString(%s) = \"%q\", want = \"%s\"", test.s, got, want)
		}
	}

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
