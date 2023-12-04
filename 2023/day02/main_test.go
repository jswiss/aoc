package main

import (
	"testing"
)

func TestFindDigitBeforeSubstring(t *testing.T) {
	line := "33 red, 18 blue"
	want := 33
	got, _ := FindDigitBeforeSubstring(line, "red")
	if got != want {
		t.Fatalf(`Yo, this ain't working. You wanted %d, but you got %d, given input %s.`, want, got, line)
	}
}

func TestCheckPossibility(t *testing.T) {
	var input []int
	input = append(input, 12, 11, 15, 17)
	want := false
	got := CheckPossibility(input, 12)
	if got != want {
		t.Fatalf(`Yo, this ain't working. You wanted %t, but you got %t.`, want, got)
	}
}
