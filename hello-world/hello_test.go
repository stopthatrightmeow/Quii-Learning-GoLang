package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Got %q Want %q", got, want)
		}
	}

	// Saying Hello, person when a name is passed
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Stopthatrightmeow", "")
		want := "Hello, Stopthatrightmeow"

		assertCorrectMessage(t, got, want)
	})
	// Default to Hello World if no name is given
	t.Run("Saying 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
	// Spanish
	t.Run("Saying hello to people in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})
	// French
	t.Run("Saying hello to people en France", func(t *testing.T) {
		got := Hello("Merideth", "French")
		want := "Bonjour, Merideth"

		assertCorrectMessage(t, got, want)
	})
	// German
	t.Run("Saying hello to people in German", func(t *testing.T) {
		got := Hello("Chloe", "German")
		want := "Hallo, Chloe"

		assertCorrectMessage(t, got, want)
	})
}
