package main

import (
	"context"
	"fmt"
	"net/http"
)

// Software often kicks off long-running, resource-intensive processes (often in goroutines). If the action that caused this gets cancelled or fails for some reason you need to stop these processes in a consistent way through your application.
// If you don't manage this your snappy Go application that you're so proud of could start having difficult to debug performance problems.
// In this chapter we'll use the package context to help us manage long-running processes.
// We're going to start with a classic example of a web server that when hit kicks off a potentially long-running process to fetch some data for it to return in the response.
// We will exercise a scenario where a user cancels the request before the data can be retrieved and we'll make sure the process is told to give up.
// I've set up some code on the happy path to get us started. Here is our server code.

// func Server(store Store) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, store.Fetch())
// 	}
// }

// The function Server takes a Store and returns us a http.HandlerFunc. Store is defined as:

// type Store interface {
// 	Fetch() string
// }

// The returned function calls the store's Fetch method to get the data and writes it to the response.

// ---------

// Now that we have a happy path, we want to make a more realistic scenario where the Store can't finish aFetch before the user cancels the request.

// Our handler will need a way of telling the Store to cancel the work so update the interface.

// type Store interface {
// 	Fetch() string
// 	Cancel()
// }

// Remember to be disciplined with TDD. Write the minimal amount of code to make our test pass.

// func Server(store Store) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		store.Cancel()
// 		fmt.Fprint(w, store.Fetch())
// 	}
// }

// This makes this test pass but it doesn't feel good does it! We surely shouldn't be cancelling Store before we fetch on every request.
// By being disciplined it highlighted a flaw in our tests, this is a good thing!

// Run both tests and the happy path test should now be failing and now we're forced to do a more sensible implementation.

// func Server(store Store) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		ctx := r.Context()

// 		data := make(chan string, 1)

// 		go func() {
// 			data <- store.Fetch()
// 		}()

// 		select {
// 		case d := <-data:
// 			fmt.Fprint(w, d)
// 		case <-ctx.Done():
// 			store.Cancel()
// 		}
// 	}
// }

// context has a method Done() which returns a channel which gets sent a signal when the context is "done" or "cancelled".
// We want to listen to that signal and call store.Cancel if we get it but we want to ignore it if our Store manages to Fetch before it.
// To manage this we run Fetch in a goroutine and it will write the result into a new channel data.
// We then use select to effectively race to the two asynchronous processes and then we either write a response or Cancel.

// ----------

// We'll have to change our existing tests as their responsibilities are changing.
// The only thing our handler is responsible for now is making sure it sends a context through to the downstream Store and that it handles the error that will come from the Store when it is cancelled.

// Let's update our Store interface to show the new responsibilities.

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// Delete the code inside our handler for now

// func Server(store Store) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 	}
// }

// We have to make our spy act like a real method that works with context.

// We are simulating a slow process where we build the result slowly by appending the string, character by character in a goroutine.
// When the goroutine finishes its work it writes the string to the data channel.
// The goroutine listens for the ctx.Done and will stop the work if a signal is sent in that channel.
// Finally the code uses another select to wait for that goroutine to finish its work or for the cancellation to occur.
// It's similar to our approach from before, we use Go's concurrency primitives to make two asynchronous processes race each other to determine what we return.
// You'll take a similar approach when writing your own functions and methods that accept a context so make sure you understand what's going on.

// Our happy path should be... happy. Now we can fix the other test.

// func Server(store Store) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		data, _ := store.Fetch(r.Context())
// 		fmt.Fprint(w, data)
// 	}
// }

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return // TODO: log error however you like
		}

		fmt.Fprint(w, data)
	}
}

// We can see after this that the server code has become simplified as it's no longer explicitly responsible for cancellation,
// it simply passes through context and relies on the downstream functions to respect any cancellations that may occur.
