package main

import (
	"reflect"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestSymbolIndices(t *testing.T) {
	idx := 0
	line := ".*.$.aaa...555"
	var symbol Symbol
	symbol.line = 0
	symbol.symbols = []int{1, 3}
	want := symbol
	got := GetSymbolIndices(idx, line)

	if reflect.DeepEqual(want, got) == false {
		t.Fatalf(`Yo, this ain't working. You wanted %v, but you got %v`, want, got)
	}

}
