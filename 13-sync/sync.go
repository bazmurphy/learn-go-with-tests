package sync_chapter

import "sync"

// We want to make a counter which is safe to use concurrently.
// We'll start with an unsafe counter and verify its behaviour works in a single-threaded environment.
// Then we'll exercise its unsafeness, with multiple goroutines trying to use the counter via a test, and fix it.

// ----------

// This should be trivial for Go experts like us. We need to keep some state for the counter in our datatype and then increment it on every Inc call

// type Counter struct {
// 	value int
// }

// func (c *Counter) Inc() {
// 	c.value++
// }

// func (c *Counter) Value() int {
// 	return c.value
// }

// ----------

// A simple solution is to add a lock to our Counter, ensuring only one goroutine can increment the counter at a time. Go's Mutex provides such a lock:
// "A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex."

// type Counter struct {
// 	mu    sync.Mutex
// 	value int
// }

// func (c *Counter) Inc() {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.value++
// }

// func (c *Counter) Value() int {
// 	return c.value
// }

// What this means is any goroutine calling Inc will acquire the lock on Counter if they are first. All the other goroutines will have to wait for it to be Unlocked before getting access.
// If you now re-run the test it should now pass because each goroutine has to wait its turn before making a change.

// Our test passes but our code is still a bit dangerous
// A look at the documentation of sync.Mutex tells us why
// "A Mutex must not be copied after first use."
// When we pass our Counter (by value) to assertCounter it will try and create a copy of the mutex.
// To solve this we should pass in a pointer to our Counter instead, so change the signature of assertCounter

// To solve this we should pass in a pointer to our Counter instead, so change the signature of assertCounter
// func assertCounter(t testing.TB, got *Counter, want int)

// Our tests will no longer compile because we are trying to pass in a Counter rather than a *Counter. To solve this I prefer to create a constructor which shows readers of your API that it would be better to not initialise the type yourself.

type Counter struct {
	mu    sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
