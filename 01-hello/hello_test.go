package hello

import "testing"

// a test takes t a pointer to the testing.T type
// we can then use the methods/fields of testing.T in our test function
func TestHello(t *testing.T) {
	// t.Run a method that allows subtests within a single test function
	// Run(name string, f func(t *testing.T)) bool
	// [1] name of the subtest (string)
	// [2] the subtest function
	// returns a result of the subtest (boolean)
	t.Run("to a person", func(t *testing.T) {
		got := Hello("Baz", "")
		want := "Hello, Baz"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", spanish)
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Lauren", french)
		want := "Bonjour, Lauren"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in German", func(t *testing.T) {
		got := Hello("Clara", german)
		want := "Hallo, Clara"
		assertCorrectMessage(t, got, want)
	})
}

// our own defined helper function
// testing.TB is the Testing&Benchmark types
func assertCorrectMessage(t testing.TB, got, want string) {
	// By calling t.Helper() at the beginning of your helper function,
	// you instruct the testing framework to adjust the recorded file name and line number
	// to match the location where the helper function was called from,
	// rather than the location of the helper function itself.
	t.Helper()
	if got != want {
		// now when t.Errorf the file name and line number will be within the test function
		// and not here in the helper function
		t.Errorf("got %q want %q", got, want)
	}
}
