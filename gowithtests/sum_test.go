package main

import "testing"

func TestSum(t *testing.T){

	s := []int{1,2,3,4,5}
	got := Sum(s)
	want := 15

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
