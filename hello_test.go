package main

import "testing"

func TestHello(t *testing.T) {
	assertMesg := func(got string, want string, t *testing.T) {
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		current := Hello("Chris")
		expected := "Hello, Chris"

		assertMesg(current, expected, t)

	})
	t.Run("say 'Hello, World' when empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertMesg(got, want, t)
	})
}
