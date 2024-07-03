package context_chapter

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// We have a corresponding spy for Store which we use in a test.

// type SpyStore struct {
// 	response string
// }

// func (s *SpyStore) Fetch() string {
// 	return s.response
// }

// func TestServer(t *testing.T) {
// 	data := "hello world"
// 	svr := Server(&SpyStore{data})

// 	request := httptest.NewRequest(http.MethodGet, "/", nil)
// 	response := httptest.NewRecorder()

// 	svr.ServeHTTP(response, request)

// 	if response.Body.String() != data {
// 		t.Errorf("got %s, want %s", response.Body.String(), data)
// 	}
// }

// We will need to adjust our spy so it takes some time to return data and a way of knowing it has been told to cancel. We'll also rename it to SpyStore as we are now observing the way it is called. It'll have to add Cancel as a method to implement the Store interface.

// type SpyStore struct {
// 	response  string
// 	cancelled bool
// }

// func (s *SpyStore) Fetch() string {
// 	time.Sleep(100 * time.Millisecond)
// 	return s.response
// }

// func (s *SpyStore) Cancel() {
// 	s.cancelled = true
// }

// Let's add a new test where we cancel the request before 100 milliseconds and check the store to see if it gets cancelled.

// func TestServer(t *testing.T) {
// 	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
// 		data := "hello world"
// 		store := &SpyStore{response: data}
// 		svr := Server(store)

// 		request := httptest.NewRequest(http.MethodGet, "/", nil)

// 		cancellingCtx, cancel := context.WithCancel(request.Context())
// 		time.AfterFunc(5*time.Millisecond, cancel)
// 		request = request.WithContext(cancellingCtx)

// 		response := httptest.NewRecorder()

// 		svr.ServeHTTP(response, request)

// 		if !store.cancelled {
// 			t.Errorf("store was not told to cancel")
// 		}
// 	})
// }

// From the Go Blog: Context
// The context package provides functions to derive new Context values from existing ones. These values form a tree: when a Context is canceled, all Contexts derived from it are also canceled.
// It's important that you derive your contexts so that cancellations are propagated throughout the call stack for a given request.
// What we do is derive a new cancellingCtx from our request which returns us a cancel function. We then schedule that function to be called in 5 milliseconds by using time.AfterFunc. Finally we use this new context in our request by calling request.WithContext.

// ---------

// We'll need to update our happy path test to assert that it does not get cancelled.

// func TestServer(t *testing.T) {
// 	t.Run("returns data from store", func(t *testing.T) {
// 		data := "hello, world"
// 		store := &SpyStore{response: data}
// 		svr := Server(store)

// 		request := httptest.NewRequest(http.MethodGet, "/", nil)
// 		response := httptest.NewRecorder()

// 		svr.ServeHTTP(response, request)

// 		if response.Body.String() != data {
// 			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
// 		}

// 		if store.cancelled {
// 			t.Error("it should not have cancelled the store")
// 		}
// 	})

// 	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
// 		data := "hello world"
// 		store := &SpyStore{response: data}
// 		svr := Server(store)

// 		request := httptest.NewRequest(http.MethodGet, "/", nil)

// 		cancellingCtx, cancel := context.WithCancel(request.Context())
// 		time.AfterFunc(5*time.Millisecond, cancel)
// 		request = request.WithContext(cancellingCtx)

// 		response := httptest.NewRecorder()

// 		svr.ServeHTTP(response, request)

// 		if !store.cancelled {
// 			t.Errorf("store was not told to cancel")
// 		}
// 	})
// }

// ----------

// We can refactor our test code a bit by making assertion methods on our spy

// type SpyStore struct {
// 	response  string
// 	cancelled bool
// 	t         *testing.T
// }

// func (s *SpyStore) Fetch() string {
// 	time.Sleep(100 * time.Millisecond)
// 	return s.response
// }

// func (s *SpyStore) Cancel() {
// 	s.cancelled = true
// }

// func (s *SpyStore) assertWasCancelled() {
// 	s.t.Helper()
// 	if !s.cancelled {
// 		s.t.Error("store was not told to cancel")
// 	}
// }

// func (s *SpyStore) assertWasNotCancelled() {
// 	s.t.Helper()
// 	if s.cancelled {
// 		s.t.Error("store was told to cancel")
// 	}
// }

// Remember to pass in the *testing.T when creating the spy.

// func TestServer(t *testing.T) {
// 	data := "hello world"

// 	t.Run("returns data from store", func(t *testing.T) {
// 		store := &SpyStore{response: data, t: t}
// 		svr := Server(store)

// 		request := httptest.NewRequest(http.MethodGet, "/", nil)
// 		response := httptest.NewRecorder()

// 		svr.ServeHTTP(response, request)

// 		if response.Body.String() != data {
// 			t.Errorf("got %s, want %s", response.Body.String(), data)
// 		}

// 		store.assertWasNotCancelled()
// 	})

// 	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
// 		store := &SpyStore{response: data, t: t}
// 		svr := Server(store)

// 		request := httptest.NewRequest(http.MethodGet, "/", nil)

// 		cancellingCtx, cancel := context.WithCancel(request.Context())
// 		time.AfterFunc(5*time.Millisecond, cancel)
// 		request = request.WithContext(cancellingCtx)

// 		response := httptest.NewRecorder()

// 		svr.ServeHTTP(response, request)

// 		store.assertWasCancelled()
// 	})
// }

// ----------

// This approach is ok, but is it idiomatic?
// Does it make sense for our web server to be concerned with manually cancelling Store?
// What if Store also happens to depend on other slow-running processes?
// We'll have to make sure that Store.Cancel correctly propagates the cancellation to all of its dependants.
// One of the main points of context is that it is a consistent way of offering cancellation.

// From the go doc

// "Incoming requests to a server should create a Context, and outgoing calls to servers should accept a Context.
// The chain of function calls between them must propagate the Context, optionally replacing it with a derived Context created using WithCancel, WithDeadline, WithTimeout, or WithValue.
// When a Context is canceled, all Contexts derived from it are also canceled."

// From the Go Blog: Context again:

// "At Google, we require that Go programmers pass a Context parameter as the first argument to every function on the call path between incoming and outgoing requests.
// This allows Go code developed by many different teams to interoperate well.
// It provides simple control over timeouts and cancelation and ensures that critical values like security credentials transit Go programs properly."

// Pause for a moment and think of the ramifications of every function having to send in a context, and the ergonomics of that.)

// Feeling a bit uneasy? Good. Let's try and follow that approach though and instead pass through the context to our Store and let it be responsible.
// That way it can also pass the context through to its dependants and they too can be responsible for stopping themselves.

// Update our SpyStore

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

// We need to test that we do not write any kind of response on the error case.
// Sadly httptest.ResponseRecorder doesn't have a way of figuring this out so we'll have to roll our own spy to test for this.

// Our SpyResponseWriter implements http.ResponseWriter so we can use it in the test.

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestServer(t *testing.T) {
	t.Run("returns data from store", func(t *testing.T) {
		data := "hello world"
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got %s, want %s", response.Body.String(), data)
		}
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello world"
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		if response.written {
			t.Error("a response should not have been written")
		}
	})
}
