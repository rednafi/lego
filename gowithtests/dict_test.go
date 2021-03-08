package main

import "testing"

//import "fmt"

func TestDict(t *testing.T) {

	t.Run("string assert", func(t *testing.T) {

		d := Dict{"test": "this is just a test"}
		got, _ := d.Search("test")
		want := "this is just a test"

		assertString(t, got, want)

	})

	t.Run("unknown key", func(t *testing.T) {

		d := Dict{"test": "this is just a test"}
		_, ok := d.Search("unknown")

		_ = "this is just a test"

		assertKeyError(t, ok)

	})

	t.Run("Add", func(t *testing.T) {

		d := Dict{"test": "this is just a test"}
		err := d.Add("bohemian", "rhapsody")

		if err != nil {
			t.Errorf("dict add failed")
		}
	})

	t.Run("Update", func(t *testing.T) {

		d := Dict{"bohemian": "rhapsody"}
		err := d.Update("bohemian", "rhapsody")

		if err != nil {
			t.Errorf("dict update failed, key does not exist")
		}
	})

	t.Run("Delete", func(t *testing.T) {

		d := Dict{"bohemian": "rhapsody"}
		err := d.Delete("bohemian")

		if err != nil {
			t.Errorf("key delete failed, key does not exist")
		}
	})

}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}

func assertKeyError(t testing.TB, ok error) {
	t.Helper()
	if ok == nil {
		t.Errorf("key error did not occur when expecting that")
	}

}
