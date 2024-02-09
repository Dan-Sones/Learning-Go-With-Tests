package main

import "testing"

//Write a test
//Make the compiler pass
//Run the test, see that it fails and check the error message is meaningful
//Write enough code to make the test pass
//Refactor

func TestHello(t *testing.T) {
	t.Run("Saying Hello to people", func(t *testing.T) {
		got := Hello("Daniel", "")
		want := "Hello, Daniel"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in french", func(t *testing.T) {
		got := Hello("Melody", "French")
		want := "Bonjour, Melody"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// Lets the test suite know this is a helper method,
	// If the test fails, the pointer won't be here, it will be in the parent function
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q ", got, want)
	}
}
