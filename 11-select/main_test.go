package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// func TestRacer(t *testing.T) {
// 	slowURL := "http://facebook.com"
// 	fastURL := "http://www.quii.dev"

// 	want := fastURL
// 	got := Racer(slowURL, fastURL)

// 	if got != want {
// 		t.Errorf("got %q, want %q", got, want)
// 	}
// }

// This may or may not make the test pass for you. The problem is we're reaching out to real websites to test our own logic.
// Testing code that uses HTTP is so common that Go has tools in the standard library to help you test it.
// In the mocking and dependency injection chapters, we covered how ideally we don't want to be relying on external services to test our code because they can be
// - Slow
// - Flaky
// - Can't test edge cases
// In the standard library, there is a package called net/http/httptest which enables users to easily create a mock HTTP server.

// func TestRacer(t *testing.T) {
// 	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		time.Sleep(20 * time.Millisecond)
// 		w.WriteHeader(http.StatusOK)
// 	}))

// 	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(http.StatusOK)
// 	}))

// 	slowURL := slowServer.URL
// 	fastURL := fastServer.URL

// 	want := fastURL
// 	got := Racer(slowURL, fastURL)

// 	if got != want {
// 		t.Errorf("got %q, want %q", got, want)
// 	}

// 	slowServer.Close()
// 	fastServer.Close()
// }

// httptest.NewServer takes an http.HandlerFunc which we are sending in via an anonymous function.
// http.HandlerFunc is a type that looks like this: type HandlerFunc func(ResponseWriter, *Request)
// All it's really saying is it needs a function that takes a ResponseWriter and a Request, which is not too surprising for an HTTP server.
// It turns out there's really no extra magic here, this is also how you would write a real HTTP server in Go. The only difference is we are wrapping it in an httptest.NewServer which makes it easier to use with testing, as it finds an open port to listen on and then you can close it when you're done with your test.

// func TestRacer(t *testing.T) {
// 	slowServer := makeDelayedServer(20 * time.Millisecond)
// 	fastServer := makeDelayedServer(0 * time.Millisecond)

// 	// By prefixing a function call with defer it will now call that function at the end of the containing function.
// 	// Sometimes you will need to clean up resources, such as closing a file or in our case closing a server so that it does not continue to listen to a port.
// 	// You want this to execute at the end of the function, but keep the instruction near where you created the server for the benefit of future readers of the code.
// 	defer slowServer.Close()
// 	defer fastServer.Close()

// 	slowURL := slowServer.URL
// 	fastURL := fastServer.URL

// 	want := fastURL
// 	got := Racer(slowURL, fastURL)

// 	if got != want {
// 		t.Errorf("got %q, want %q", got, want)
// 	}
// }

// func makeDelayedServer(delay time.Duration) *httptest.Server {
// 	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		time.Sleep(delay)
// 		w.WriteHeader(http.StatusOK)
// 	}))
// }

// func TestRacer(t *testing.T) {
// 	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
// 		slowServer := makeDelayedServer(20 * time.Millisecond)
// 		fastServer := makeDelayedServer(0 * time.Millisecond)

// 		defer slowServer.Close()
// 		defer fastServer.Close()

// 		slowURL := slowServer.URL
// 		fastURL := fastServer.URL

// 		want := fastURL
// 		got, err := Racer(slowURL, fastURL)

// 		if err != nil {
// 			t.Error("expected no error but got one")
// 		}

// 		if got != want {
// 			t.Errorf("got %q, want %q", got, want)
// 		}
// 	})

// 	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
// 		serverA := makeDelayedServer(11 * time.Second)
// 		serverB := makeDelayedServer(12 * time.Second)

// 		defer serverA.Close()
// 		defer serverB.Close()

// 		_, err := Racer(serverA.URL, serverB.URL)

// 		if err == nil {
// 			t.Error("expected an error but didn't get one")
// 		}
// 	})
// }

// func makeDelayedServer(delay time.Duration) *httptest.Server {
// 	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		time.Sleep(delay)
// 		w.WriteHeader(http.StatusOK)
// 	}))
// }

// We've made our test servers take longer than 10s to return to exercise this scenario and we are expecting Racer to return two values now, the winning URL (which we ignore in this test with _) and an error.

// ----------

// The problem we have is that this test takes 10 seconds to run. For such a simple bit of logic, this doesn't feel great.

// ----------

func TestRacer(t *testing.T) {
	t.Run("compares speed of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableReacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected and error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

// Our users and our first test can use Racer (which uses ConfigurableRacer under the hood) and our sad path test can use ConfigurableRacer.
