package arrays_and_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers (array)", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	// Every test has a cost
	// Having too many tests can turn in to a real problem and it just adds more overhead in maintenance
	// having two tests for this function is redundant
	// If it works for a slice of one size it's very likely it'll work for a slice of any size

	// t.Run("collection of any size (slice)", func(t *testing.T) {
	// 	numbers := []int{1, 2, 3}

	// 	got := Sum(numbers)
	// 	want := 6

	// 	if got != want {
	// 		t.Errorf("got %d want %d given, %v", got, want, numbers)
	// 	}
	// })
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	// if got != want {
	// 	t.Errorf("got %v want %v", got, want)
	// }

	// Go does not let you use equality operators with slices
	// you could iterate over each got and want slice and check their values
	// but we can use reflect.DeepEqual which checks if any two variables are the same
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	// we can assign a function to a variable
	// It's not shown here, but this technique can be useful when you want to bind a function to other local variables in "scope" (e.g between some {}).
	// It also allows you to reduce the surface area of your API.
	// By defining this function inside the test, it cannot be used by other functions in this package.
	// Hiding variables and functions that don't need to be exported is an important design consideration.
	// A handy side-effect of this is this adds a little type-safety to our code.
	// If a developer mistakenly adds a new test with checkSums(t, got, "dave") the compiler will stop them in their tracks.
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		// if !reflect.DeepEqual(got, want) {
		// 	t.Errorf("got %v want %v", got, want)
		// }
		checkSums(t, got, want)
	})

	// Compile time errors are our friend because they help us write software that works,
	// runtime errors are our enemies because they affect our users.
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		// if !reflect.DeepEqual(got, want) {
		// 	t.Errorf("got %v want %v", got, want)
		// }
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}
