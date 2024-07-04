package main

import (
	"sync"
	"testing"
)

// We want to make a counter which is safe to use concurrently.
// We'll start with an unsafe counter and verify its behaviour works in a single-threaded environment.
// Then we'll exercise its unsafeness, with multiple goroutines trying to use the counter via a test, and fix it.

// func TestCounter(t *testing.T) {
// 	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
// 		counter := Counter{}
// 		counter.Inc()
// 		counter.Inc()
// 		counter.Inc()

// 		if counter.Value() != 3 {
// 			t.Errorf("got %d, want %d", counter.Value(), 3)
// 		}
// 	})
// }

// There's not a lot to refactor but given we're going to write more tests around Counter we'll write a small assertion function assertCount so the test reads a bit clearer.

// func TestCounter(t *testing.T) {
// 	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
// 		counter := Counter{}
// 		counter.Inc()
// 		counter.Inc()
// 		counter.Inc()

// 		assertCounter(t, counter, 3)
// 	})
// }

// func assertCounter(t testing.TB, got Counter, want int) {
// 	t.Helper()
// 	if got.Value() != want {
// 		t.Errorf("got %d, want %d", got.Value(), want)
// 	}
// }

// func TestCounter(t *testing.T) {
// 	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
// 		counter := Counter{}
// 		counter.Inc()
// 		counter.Inc()
// 		counter.Inc()

// 		assertCounter(t, counter, 3)
// 	})

// 	// That was easy enough but now we have a requirement that it must be safe to use in a concurrent environment. We will need to write a failing test to exercise this.

// 	t.Run("it runs safely concurrently", func(t *testing.T) {
// 		wantedCount := 1000
// 		counter := Counter{}

// 		var wg sync.WaitGroup
// 		wg.Add(wantedCount)

// 		for i := 0; i < wantedCount; i++ {
// 			go func() {
// 				counter.Inc()
// 				wg.Done()
// 			}()
// 		}
// 		wg.Wait()

// 		assertCounter(t, counter, wantedCount)
// 	})
// }

// func assertCounter(t testing.TB, got Counter, want int) {
// 	t.Helper()
// 	if got.Value() != want {
// 		t.Errorf("got %d, want %d", got.Value(), want)
// 	}
// }

// This will loop through our wantedCount and fire a goroutine to call counter.Inc().
// We are using sync.WaitGroup which is a convenient way of synchronising concurrent processes.
// "A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished."
// By waiting for wg.Wait() to finish before making our assertions we can be sure all of our goroutines have attempted to Inc the Counter.

// The test will probably fail with a different number, but nonetheless it demonstrates it does not work when multiple goroutines are trying to mutate the value of the counter at the same time.

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
