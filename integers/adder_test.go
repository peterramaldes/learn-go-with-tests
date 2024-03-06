package integers

import (
	"fmt"
	"testing"
)

func ExampleAdder() {
	fmt.Println(Add(1, 6))
	// Output: 7
}

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
