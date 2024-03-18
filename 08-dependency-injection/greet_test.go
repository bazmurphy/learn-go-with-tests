package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// The Buffer type from the bytes package implements the Writer interface,
	// because it has the method Write(p []byte) (n int, err error).

	// The bytes.Buffer type is an implementation of the io.Writer interface in Go's standard library.
	// It provides a way to write data to an in-memory buffer, which can be useful for various purposes,
	// such as
	// - Testing: it allows you to capture the output of a function that writes to an io.Writer interface, instead of writing to stdout. This makes it easier to inspect and verify the output in tests.
	// - Building Strings: You can write data to a bytes.Buffer and then convert it to a string using the String() method. This can be more efficient than concatenating strings in certain situations.
	// - Buffering Data: The bytes.Buffer can be used to buffer data before writing it to a larger destination, like a file or network connection. This can improve performance by reducing the number of write operations.

	// Create a new Buffer value that implements the Writer interface.
	// This allows us to write the output of the Greet function to the buffer instead of stdout (standard output).
	// This is useful for testing purposes, as it allows us to capture and inspect the output of the function.
	buffer := bytes.Buffer{}
	// By creating a new bytes.Buffer instance with bytes.Buffer{},
	// we're initializing an empty buffer that we can write data to using the Write method
	// or other related methods provided by the bytes.Buffer type.

	Greet(&buffer, "Baz")

	got := buffer.String()

	want := "Hello, Baz"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
