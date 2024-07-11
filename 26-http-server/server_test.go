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

// type StubPlayerStore struct {
// 	scores map[string]int
// }

// func (s *StubPlayerStore) GetPlayerScore(name string) int {
// 	score := s.scores[name]
// 	return score
// }

// A map is a quick and easy way of making a stub key/value store for our tests.
// Now let's create one of these stores for our tests and send it into our PlayerServer.

// func TestGETPlayers(t *testing.T) {
// 	store := StubPlayerStore{
// 		map[string]int{
// 			"Pepper": 20,
// 			"Floyd":  10,
// 		},
// 	}

// 	server := &PlayerServer{&store}

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

// Our tests now pass and are looking better.
// The intent behind our code is clearer now due to the introduction of the store.
// We're telling the reader that because we have this data in a PlayerStore that when you use it with a PlayerServer you should get the following responses.

// -----------

// Add a missing player scenario to our existing suite

// func TestGETPlayers(t *testing.T) {
// 	store := StubPlayerStore{
// 		map[string]int{
// 			"Pepper": 20,
// 			"Floyd":  10,
// 		},
// 	}

// 	server := &PlayerServer{&store}

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

// 	t.Run("returns 404 on missing players", func(t *testing.T) {
// 		request := newGetScoreRequest("Apollo")
// 		response := httptest.NewRecorder()

// 		server.ServeHTTP(response, request)

// 		got := response.Code
// 		want := http.StatusNotFound

// 		if got != want {
// 			t.Errorf("got status %d want %d", got, want)
// 		}
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

// Update the other two tests to assert on the status and fix the code.

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil, // (!) added much later to satisfy additions below
	}

	server := &PlayerServer{&store}

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
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

// We're checking the status in all our tests now so I made a helper assertStatus to facilitate that.

// Now our first two tests fail because of the 404 instead of 200, so we can fix PlayerServer to only return not found if the score is 0.

// -----------

// Now that we can retrieve scores from a store it now makes sense to be able to store new scores.

// For a start let's just check we get the correct status code if we hit the particular route with POST.
// This lets us drive out the functionality of accepting a different kind of request and handling it differently to GET /players/{name}.
// Once this works we can then start asserting on our handler's interaction with the store.

// func TestStoreWins(t *testing.T) {
// 	store := StubPlayerStore{
// 		map[string]int{},
// 	}
// 	server := &PlayerServer{&store}

// 	t.Run("it returns accepted on POST", func(t *testing.T) {
// 		request, _ := http.NewRequest(http.MethodPost, "/players.Pepper", nil)
// 		response := httptest.NewRecorder()

// 		server.ServeHTTP(response, request)

// 		assertStatus(t, response.Code, http.StatusAccepted)
// 	})
// }

// -----------

// Next, we want to check that when we do our POST /players/{name} that our PlayerStore is told to record the win.

// We can accomplish this by extending our StubPlayerStore with a new RecordWin method and then spy on its invocations.

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

// Now extend our test to check the number of invocations for a start

// -----------

// We need to update our code where we create a StubPlayerStore as we've added a new field

// func TestStoreWins(t *testing.T) {
// 	store := StubPlayerStore{
// 		map[string]int{},
// 		nil,
// 	}
// 	server := &PlayerServer{&store}

// 	t.Run("it records wins when POST", func(t *testing.T) {
// 		request := newPostWinRequest("Pepper")
// 		response := httptest.NewRecorder()

// 		server.ServeHTTP(response, request)

// 		assertStatus(t, response.Code, http.StatusAccepted)

// 		if len(store.winCalls) != 1 {
// 			t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
// 		}
// 	})
// }

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

// As we're only asserting the number of calls rather than the specific values it makes our initial iteration a little smaller.

// -----------

// Run the tests and it should be passing! Obviously "Bob" isn't exactly what we want to send to RecordWin, so let's further refine the test.

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
	}
	server := &PlayerServer{&store}

	t.Run("it records wins on POST", func(t *testing.T) {
		player := "Pepper"

		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], player)
		}
	})
}

// Now that we know there is one element in our winCalls slice we can safely reference the first one and check it is equal to player.
