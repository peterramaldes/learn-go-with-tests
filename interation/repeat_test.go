package interation

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	fmt.Println(Repeat("b", 10))
	// Output: bbbbbbbbbb
}

func TestRepeat(t *testing.T) {
	got := Repeat("a", 3)
	want := "aaa"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
