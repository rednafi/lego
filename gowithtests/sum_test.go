package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("collections of 5 numbers", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		got := Sum(s)
		want := 15
		assertCorrectMessage(t, got, want)

	})

	t.Run("collection of any size", func(t *testing.T) {

		s := []int{1, 2, 3}
		got := Sum(s)
		want := 6
		assertCorrectMessage(t, got, want)

	})

}

func TestSumAll(t *testing.T) {
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("sum where one input is empty slice", func(t *testing.T) {
		got := SumAll([]int{}, []int{0, 9})
		want :=  []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
