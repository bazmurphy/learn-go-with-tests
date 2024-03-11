package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("b", 10)
	fmt.Println(repeated)
	// Output: bbbbbbbbbb
}

// Go has built in support for Benchmarking
// *testing.B for Benchmarks
func BenchmarkRepeat(b *testing.B) {
	// it runs b.N times and measures how long it takes
	// the framework itself decides how many times to run the loop
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

// X ns/op means the function takes on average X nanoseconds to run (on this specific computer)
