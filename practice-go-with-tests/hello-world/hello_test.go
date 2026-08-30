package main

import "testing"

func TestHello(t *testing.T) {
	expect := "Hello, Benny"

	result := Hello("Benny")
	if result != expect {
		t.Errorf(`call to Hello() = %q, want %q`, result, expect)
	}

}
