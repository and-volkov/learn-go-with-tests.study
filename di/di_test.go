package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Andrei")

	got := buffer.String()
	want := "Hello, Andrei"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
