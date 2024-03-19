package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Peter")

	got := buffer.String()
	want := "Hello, Peter"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
