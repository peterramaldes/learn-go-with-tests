package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Peter")
		want := "Hello, Peter"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("saying 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func ExampleHello() {
	fmt.Println(Hello("Peter"))
	// Output:
	// Hello, Peter
}
