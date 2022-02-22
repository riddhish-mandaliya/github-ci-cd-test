package hello

import (
	"testing"
)

func TestMain(t *testing.T) {
	want := "Hello world"
	got := PrefixHello("world")
	if got != want {
		t.Errorf("Expecting %s, got %s", want, got)
	}
}
