package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hello stranger!"
	if got := Greet(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
