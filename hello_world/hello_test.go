package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Sammy")
	want := "Hello, Sammy"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
