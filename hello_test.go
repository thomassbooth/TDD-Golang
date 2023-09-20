package main

import "testing"

//Tests need to be in a file with a name like xxx_test.go
//Test functions must start with the word Test
//Test functions only take 1 Parameter: t *testing.T

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "hello Chris"

		assertCorrectMessage(t, got, want)
	})
	t.Run("Say, hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "hello world"

		assertCorrectMessage(t, got, want)

	})

	t.Run("In spanish!", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "hola Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In french!", func(t *testing.T) {
		got := Hello("Dave", "French")
		want := "bonjour Dave"

		assertCorrectMessage(t, got, want)
	})
}

// For helper functions, it's a good idea to accept a testing.TB
// which is an interface that *testing.T and *testing.B
func assertCorrectMessage(t testing.TB, got, want string) {

	//t.Helper() tells the test suite that this method is a helper
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
