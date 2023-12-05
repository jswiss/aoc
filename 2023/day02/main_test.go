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
	input := []int{12, 11, 15, 17}
	want := false
	got := CheckPossibility(input, 12)
	if got != want {
		t.Fatalf(`Yo, this ain't working. You wanted %t, but you got %t.`, want, got)
	}
}

func TestGetMax(t *testing.T) {
	input := []int{65, 4, 88, 23, 6}
	want := 88
	got := GetMax(input)
	if got != want {
		t.Fatalf(`Yo, this ain't working. You wanted %d, but you got %d.`, want, got)
	}
}
