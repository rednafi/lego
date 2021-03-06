package main

import "testing"

// Here testing.T is for unittesting and testing.B is for benchmarking
func TestAdder(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("add two small integers", func(t *testing.T) {
		got := Adder(1, 2)
		want := 3

		assertCorrectMessage(t, got, want)

	})
}
