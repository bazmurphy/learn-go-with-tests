package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// Go has a concept of Examples in tests
// Examples are compiled (and optionally executed) as part of a package's test suite
// As with typical tests, examples are functions that reside in a package's _test.go files
// (!) If your code changes so that the example is no longer valid, your build will fail.
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
