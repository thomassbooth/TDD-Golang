package main

import "testing"

//Tests need to be in a file with a name like xxx_test.go
//Test functions must start with the word Test
//Test functions only take 1 Parameter: t *testing.T

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "hello Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
