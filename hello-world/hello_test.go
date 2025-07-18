package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Andrei", "")
		want := "Hello, Andrei"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Mark", "French")
		want := "Bonjour, Mark"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Russian", func(t *testing.T) {
		got := Hello("Ivan", "Russian")
		want := "Привет, Ivan"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, world' when empty string provided", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
