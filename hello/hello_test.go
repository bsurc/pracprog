package main

import(
	"testing"
)
func TestHello(t *testing.T) {
	if got, want := hello("kyle"), "Hello, kyle!"; got != want {
		t.Errorf("hello failed, got: %s, want: %s", got, want)
	}
}