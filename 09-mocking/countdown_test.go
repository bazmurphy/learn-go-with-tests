package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountDown(t *testing.T) {

	// 	t.Run("prints 3 to Go!", func(t *testing.T) {
	// 		buffer := &bytes.Buffer{}

	// 		// Update the tests to inject a dependency on our Spy and assert that the sleep has been called 3 times.
	// 		SpySleeper := &SpySleeper{}

	// 		Countdown(buffer, SpySleeper)

	// 		got := buffer.String()

	// 		// The backtick syntax is another way of creating a string but lets you include things like newlines, which is perfect for our test.
	// 		want := `3
	// 2
	// 1
	// Go!`

	// 		if got != want {
	// 			t.Errorf("got %q want %q", got, want)
	// 		}
	// 	})

	// [10]
	// we can now refactor our test so one is testing what is being printed and the other one is ensuring we're sleeping between the prints.
	// Finally, we can delete our first spy as it's not used anymore.

	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		Countdown(buffer, &SpyCountdownOperations{})

		got := buffer.String()

		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}

		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

// [13]
// With our spy in place, we can create a new test for the configurable sleeper.
func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}

	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}

	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
