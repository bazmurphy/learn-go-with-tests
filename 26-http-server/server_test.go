package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Let's create a file named server_test.go and write a test for a function PlayerServer that takes in those two arguments.
// The request sent in will be to get a player's score, which we expect to be "20".

// func TestGETPlayers(t *testing.T) {
// 	t.Run("returns Pepper's score", func(t *testing.T) {
// 		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
// 		response := httptest.NewRecorder()

// 		PlayerServer(response, request)

// 		got := response.Body.String()
// 		want := "20"

// 		if got != want {
// 			t.Errorf("got %q, want %q", got, want)
// 		}
// 	})
// }

// In order to test our server, we will need a Request to send in and we'll want to spy on what our handler writes to the ResponseWriter.

// We use http.NewRequest to create a request.
// The first argument is the request's method and the second is the request's path.
// The nil argument refers to the request's body, which we don't need to set in this case.

// net/http/httptest has a spy already made for us called ResponseRecorder so we can use that.
// It has many helpful methods to inspect what has been written as a response.

// ----------

// We'll add another subtest to our suite which tries to get the score of a different player, which will break our hard-coded approach.

// func TestGETPlayers(t *testing.T) {
// 	t.Run("returns Pepper's score", func(t *testing.T) {
// 		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
// 		response := httptest.NewRecorder()

// 		PlayerServer(response, request)

// 		got := response.Body.String()
// 		want := "20"

// 		if got != want {
// 			t.Errorf("got %q, want %q", got, want)
// 		}
// 	})

// 	t.Run("returns Floyd's score", func(t *testing.T) {
// 		request, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
// 		response := httptest.NewRecorder()

// 		PlayerServer(response, request)

// 		got := response.Body.String()
// 		want := "10"

// 		if got != want {
// 			t.Errorf("got %q, want %q", got, want)
// 		}
// 	})
// }

// You may have been thinking
// Surely we need some kind of concept of storage to control which player gets what score.
// It's weird that the values seem so arbitrary in our tests.
// Remember we are just trying to take as small as steps as reasonably possible, so we're just trying to break the constant for now.

// -----------

// And we can DRY up some of the code in the tests by making some helpers

// func TestGETPlayers(t *testing.T) {
// 	t.Run("returns Pepper's score", func(t *testing.T) {
// 		request := newGetScoreRequest("Pepper")
// 		response := httptest.NewRecorder()

// 		PlayerServer(response, request)

// 		assertResponseBody(t, response.Body.String(), "20")
// 	})

// 	t.Run("returns Floyd's score", func(t *testing.T) {
// 		request := newGetScoreRequest("Floyd")
// 		response := httptest.NewRecorder()

// 		PlayerServer(response, request)

// 		assertResponseBody(t, response.Body.String(), "10")
// 	})
// }

// func newGetScoreRequest(name string) *http.Request {
// 	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
// 	return req
// }

// func assertResponseBody(t testing.TB, got, want string) {
// 	t.Helper()
// 	if got != want {
// 		t.Errorf("response body is wrong, got %q want %q", got, want)
// 	}
// }

// ---------

// We need to change our tests to instead create a new instance of our PlayerServer and then call its method ServeHTTP.

// func TestGETPlayers(t *testing.T) {
// 	server := &PlayerServer{}

// 	t.Run("returns Pepper's score", func(t *testing.T) {
// 		request := newGetScoreRequest("Pepper")
// 		response := httptest.NewRecorder()

// 		server.ServeHTTP(response, request)

// 		assertResponseBody(t, response.Body.String(), "20")
// 	})

// 	t.Run("returns Floyd's score", func(t *testing.T) {
// 		request := newGetScoreRequest("Floyd")
// 		response := httptest.NewRecorder()

// 		server.ServeHTTP(response, request)

// 		assertResponseBody(t, response.Body.String(), "10")
// 	})
// }

// func newGetScoreRequest(name string) *http.Request {
// 	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
// 	return req
// }

// func assertResponseBody(t testing.TB, got, want string) {
// 	t.Helper()
// 	if got != want {
// 		t.Errorf("response body is wrong, got %q want %q", got, want)
// 	}
// }

// ----------

// Finally, everything is compiling but the tests are failing
// This is because we have not passed in a PlayerStore in our tests. We'll need to make a stub one up.

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

// A map is a quick and easy way of making a stub key/value store for our tests.
// Now let's create one of these stores for our tests and send it into our PlayerServer.

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}

	server := &PlayerServer{&store}

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "10")
	})
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

// Our tests now pass and are looking better.
// The intent behind our code is clearer now due to the introduction of the store.
// We're telling the reader that because we have this data in a PlayerStore that when you use it with a PlayerServer you should get the following responses.

// -----------
