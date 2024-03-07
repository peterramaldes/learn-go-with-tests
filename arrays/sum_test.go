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

func TestSumAll(t *testing.T) {
	n1 := []int{17, 28}
	n2 := []int{69, 96}

	got := SumAll(n1, n2)
	want := []int{45, 165}
	if !reflect.DeepEqual(got, want) {
		arr := [][]int{n1, n2}
		t.Errorf("got '%v' want '%v', array '%v'", got, want, arr)
	}
}
