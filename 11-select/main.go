// [0] You have been asked to make a function called WebsiteRacer which takes two URLs and "races" them by hitting them with an HTTP GET and returning the URL which returned first. If none of them return within 10 seconds then it should return an error.

package main

import (
	"fmt"
	"net/http"
	"time"
)

// func Racer(a, b string) (winner string) {
// 	startA := time.Now()
// 	http.Get(a)
// 	aDuration := time.Since(startA)

// 	startB := time.Now()
// 	http.Get(b)
// 	bDuration := time.Since(startB)

// 	if aDuration < bDuration {
// 		return a
// 	}
// 	return b
// }

// This may or may not make the test pass for you. The problem is we're reaching out to real websites to test our own logic.
// Testing code that uses HTTP is so common that Go has tools in the standard library to help you test it.
// In the mocking and dependency injection chapters, we covered how ideally we don't want to be relying on external services to test our code because they can be
// - Slow
// - Flaky
// - Can't test edge cases
// In the standard library, there is a package called net/http/httptest which enables users to easily create a mock HTTP server.

// func Racer(a, b string) (winner string) {
// 	aDuration := measureResponseTime(a)
// 	bDuration := measureResponseTime(b)

// 	if aDuration < bDuration {
// 		return a
// 	}

// 	return b
// }

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }

// Why are we testing the speeds of the websites one after another when Go is great at concurrency? We should be able to check both at the same time.
// We don't really care about the exact response times of the requests, we just want to know which one comes back first.
//  To do this, we're going to introduce a new construct called select which helps us synchronise processes really easily and clearly.

// func Racer(a, b string) (winner string) {
// 	select {
// 	// select allows you to wait on multiple channels. The first one to send a value "wins" and the code underneath the case is executed.
// 	// We use ping in our select to set up two channels, one for each of our URLs. Whichever one writes to its channel first will have its code executed in the select, which results in its URL being returned (and being the winner).
// 	case <-ping(a):
// 		return a
// 	case <-ping(b):
// 		return b
// 	}
// }

// func ping(url string) chan struct{} {
// 	ch := make(chan struct{})
// 	go func() {
// 		http.Get(url)
// 		close(ch)
// 	}()
// 	return ch
// }

// We have defined a function ping which creates a chan struct{} and returns it.
// In our case, we don't care what type is sent to the channel, we just want to signal we are done and closing the channel works perfectly!
// Why struct{} and not another type like a bool? Well, a chan struct{} is the smallest data type available from a memory perspective so we get no allocation versus a bool. Since we are closing and not sending anything on the chan, why allocate anything?
// Inside the same function, we start a goroutine which will send a signal into that channel once we have completed http.Get(url).

// Always make channels
// Notice how we have to use make when creating a channel; rather than say var ch chan struct{}. When you use var the variable will be initialised with the "zero" value of the type. So for string it is "", int it is 0, etc.
// For channels the zero value is nil and if you try and send to it with <- it will block forever because you cannot send to nil channels

// select
// You'll recall from the concurrency chapter that you can wait for values to be sent to a channel with myVar := <-ch. This is a blocking call, as you're waiting for a value.
// select allows you to wait on multiple channels. The first one to send a value "wins" and the code underneath the case is executed.
// We use ping in our select to set up two channels, one for each of our URLs. Whichever one writes to its channel first will have its code executed in the select, which results in its URL being returned (and being the winner).
// After these changes, the intent behind our code is very clear and the implementation is actually simpler.

// func Racer(a, b string) (winner string, error error) {
// 	select {
// 	case <-ping(a):
// 		return a, nil
// 	case <-ping(b):
// 		return b, nil
// 	case <-time.After(10 * time.Second):
// 		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
// 	}
// }

// func ping(url string) chan struct{} {
// 	ch := make(chan struct{})
// 	go func() {
// 		http.Get(url)
// 		close(ch)
// 	}()
// 	return ch
// }

// time.After is a very handy function when using select. Although it didn't happen in our case you can potentially write code that blocks forever if the channels you're listening on never return a value. time.After returns a chan (like ping) and will send a signal down it after the amount of time you define.
// For us this is perfect; if a or b manage to return they win, but if we get to 10 seconds then our time.After will send a signal and we'll return an error.

// func Racer(a, b string, timeout time.Duration) (winner string, error error) {
// 	select {
// 	case <-ping(a):
// 		return a, nil
// 	case <-ping(b):
// 		return b, nil
// 	case <-time.After(timeout):
// 		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
// 	}
// }

// func ping(url string) chan struct{} {
// 	ch := make(chan struct{})
// 	go func() {
// 		http.Get(url)
// 		close(ch)
// 	}()
// 	return ch
// }

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableReacer(a, b, tenSecondTimeout)
}

func ConfigurableReacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

// Before rushing in to add this default value to both our tests let's listen to them.
// Do we care about the timeout in the "happy" test?
// The requirements were explicit about the timeout.
// Given this knowledge, let's do a little refactoring to be sympathetic to both our tests and the users of our code.

// What we can do is make the timeout configurable. So in our test, we can have a very short timeout and then when the code is used in the real world it can be set to 10 seconds.
