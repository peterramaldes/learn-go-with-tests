package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got '%d' want '%d', numbers '%v'", got, want, numbers)
	}
}

func TestSumTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	}

	t.Run("sum all the tails", func(t *testing.T) {
		n1 := []int{1, 2, 3, 4}
		n2 := []int{5, 6, 7, 8, 9}

		got := SumAllTails(n1, n2)
		want := []int{9, 30}

		checkSums(t, got, want)
	})

	t.Run("passing an empty array", func(t *testing.T) {
		n1 := []int{}
		n2 := []int{3, 4, 5}

		got := SumAllTails(n1, n2)
		want := []int{0, 9}

		checkSums(t, got, want)
	})

}
