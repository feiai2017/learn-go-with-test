package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q wang %q", got, want)
	}
}
