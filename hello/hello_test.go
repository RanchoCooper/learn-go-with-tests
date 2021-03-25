package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	assertResult := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Rancho", "")
		want := "Hello, Rancho"

		assertResult(t, got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertResult(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Someone", "Spanish")
		want := "Hola, Someone"

		assertResult(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Someone", "French")
		want := "Bonjour, Someone"

		assertResult(t, got, want)
	})
}
