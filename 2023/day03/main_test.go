package main

import (
	"reflect"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestSymbolIndices(t *testing.T) {
	line := ".*.$.aaa...555"
	want := []int{1, 3}
	got := GetSymbolIndices(line)

	if reflect.DeepEqual(want, got) == false {
		t.Fatalf(`Yo, this ain't working. You wanted %v, but you got %v`, want, got)
	}

}
