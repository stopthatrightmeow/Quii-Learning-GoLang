package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Stopthatrightmeow")
		want := "Hello, Stopthatrightmeow"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Saying 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
