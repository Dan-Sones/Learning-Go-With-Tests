package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Daniel")
	want := "Hello, Daniel"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
