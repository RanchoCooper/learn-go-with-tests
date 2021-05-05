package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "rancho")

	got := buffer.String()
	want := "Hello, rancho"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
