package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Paul")
	want := "Hello, Paul"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
