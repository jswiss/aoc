package main

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestReplace(t *testing.T) {
	line := "fiverckdmgbfthjdqpmpgdkmcklqtqc365mqtwo"
	want := "5rckdmgbfthjdqpmpgdkmcklqtqc365mq2"
	got := ReplaceSubstrings(line)
	if got != want {
		t.Fatalf(`Yo, this ain't working. You wanted %s, but you got %s`, want, got)
	}
}
func TestReplaceEdge(t *testing.T) {
	line := "fiverckdmgbfthjdqpmpgdkmcklqtqc365mqtwone"
	want := "5rckdmgbfthjdqpmpgdkmcklqtqc365mq2"
	got := ReplaceSubstrings(line)
	if got != want {
		t.Fatalf(`Yo, this ain't working. You wanted %s, but you got %s`, want, got)
	}
}
