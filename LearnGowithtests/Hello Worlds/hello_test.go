package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("wanted %q got %q", want, got)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Greeting the world if name is missing", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In French", func(t *testing.T) {
		got := Hello("Franc", "French")
		want := "Bonjour, Franc"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In German", func(t *testing.T) {
		got := Hello("Nando", "German")
		want := "Hallo, Nando"
		assertCorrectMessage(t, got, want)
	})
}
