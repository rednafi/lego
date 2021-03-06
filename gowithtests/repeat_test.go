package main

import "testing"

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("test iteration", func(t *testing.T) {
		got := Repeat("a")
		want := "aaaaa"
		assertCorrectMessage(t, got, want)
	})
}


func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a")
    }
}
