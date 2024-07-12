package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// type StubPlayerStore struct {
// 	scores   map[string]int
// 	winCalls []string
// }

// func (s *StubPlayerStore) GetPlayerScore(name string) int {
// 	score := s.scores[name]
// 	return score
// }

// func (s *StubPlayerStore) RecordWin(name string) {
// 	s.winCalls = append(s.winCalls, name)
// }

// func TestGETPlayers(t *testing.T) {
// 	store := StubPlayerStore{
// 		map[string]int{
// 			"Pepper": 20,
// 			"Floyd":  10,
// 		},
// 		nil,
// 	}

// 	server := NewPlayerServer(&store)

// 	t.Run("returns Pepper's score", func(t *testing.T) {
// 		request := newGetScoreRequest("Pepper")
// 		response := httptest.NewRecorder()

// 		server.ServeHTTP(response, request)

// 		assertStatus(t, response.Code, http.StatusOK)
// 		assertResponseBody(t, response.Body.String(), "20")
// 	})

// 	t.Run("returns Floyd's score", func(t *testing.T) {
// 		request := newGetScoreRequest("Floyd")
// 		response := httptest.NewRecorder()

// 		server.ServeHTTP(response, request)

// 		assertStatus(t, response.Code, http.StatusOK)
// 		assertResponseBody(t, response.Body.String(), "10")
// 	})

// 	t.Run("returns 404 on missing players", func(t *testing.T) {
// 		request := newGetScoreRequest("Apollo")
// 		response := httptest.NewRecorder()

// 		server.ServeHTTP(response, request)

// 		assertStatus(t, response.Code, http.StatusNotFound)
// 	})
// }

// func newGetScoreRequest(name string) *http.Request {
// 	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
// 	return req
// }

// func assertStatus(t testing.TB, got, want int) {
// 	t.Helper()
// 	if got != want {
// 		t.Errorf("did not get correct status, got %d, want %d", got, want)
// 	}
// }

// func assertResponseBody(t testing.TB, got, want string) {
// 	t.Helper()
// 	if got != want {
// 		t.Errorf("response body is wrong, got %q want %q", got, want)
// 	}
// }

// func TestStoreWins(t *testing.T) {
// 	store := StubPlayerStore{
// 		map[string]int{},
// 		nil,
// 	}
// 	server := &PlayerServer{&store}

// 	t.Run("it records wins on POST", func(t *testing.T) {
// 		player := "Pepper"

// 		request := newPostWinRequest(player)
// 		response := httptest.NewRecorder()

// 		server.ServeHTTP(response, request)

// 		assertStatus(t, response.Code, http.StatusAccepted)

// 		if len(store.winCalls) != 1 {
// 			t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
// 		}

// 		if store.winCalls[0] != player {
// 			t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], player)
// 		}
// 	})
// }

// func newPostWinRequest(name string) *http.Request {
// 	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
// 	return req
// }

// We'll extend the existing suite as we have some useful test functions and a fake PlayerStore to use.

// func TestLeague(t *testing.T) {
// 	store := StubPlayerStore{}
// 	server := &PlayerServer{&store}

// 	t.Run("it returns 200 on /league", func(t *testing.T) {
// 		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
// 		response := httptest.NewRecorder()

// 		server.ServeHTTP(response, request)

// 		assertStatus(t, response.Code, http.StatusOK)
// 	})
// }

// Before worrying about actual scores and JSON we will try and keep the changes small with the plan to iterate toward our goal.
// The simplest start is to check we can hit /league and get an OK back.

// ----------

// type StubPlayerStore struct {
// 	scores   map[string]int
// 	winCalls []string
// }

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
		nil, // (!) added later to satisfy below additions
	}

	server := NewPlayerServer(&store)

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

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
		nil, // (!) added later to satisfy below additions
	}

	server := NewPlayerServer(&store)

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

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

// ----------

// We'll start by trying to parse the response into something meaningful.

// func TestLeague(t *testing.T) {
// 	store := StubPlayerStore{}
// 	server := NewPlayerServer(&store)

// 	t.Run("it returns 200 on /league", func(t *testing.T) {
// 		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
// 		response := httptest.NewRecorder()

// 		server.ServeHTTP(response, request)

// 		var got []Player

// 		err := json.NewDecoder(response.Body).Decode(&got)

// 		if err != nil {
// 			t.Fatalf("Unable to parse response from server %q into slice of Player %v", response.Body, err)
// 		}

// 		assertStatus(t, response.Code, http.StatusOK)
// 	})
// }

// Why not test the JSON string?

// You could argue a simpler initial step would be just to assert that the response body has a particular JSON string.

// In my experience tests that assert against JSON strings have the following problems.

// - Brittleness. If you change the data-model your tests will fail.

// - Hard to debug. It can be tricky to understand what the actual problem is when comparing two JSON strings.

// - Poor intention. Whilst the output should be JSON, what's really important is exactly what the data is, rather than how it's encoded.

// - Re-testing the standard library. There is no need to test how the standard library outputs JSON, it is already tested. Don't test other people's code.

// Instead, we should look to parse the JSON into data structures that are relevant for us to test with.

// Data modelling

// Given the JSON data model, it looks like we need an array of Player with some fields so we have created a new type to capture this.

// type Player struct {
// 	Name string
// 	Wins int
// }

// JSON decoding

// var got []Player
// err := json.NewDecoder(response.Body).Decode(&got)

// To parse JSON into our data model we create a Decoder from encoding/json package and then call its Decode method.
// To create a Decoder it needs an io.Reader to read from which in our case is our response spy's Body.

// Decode takes the address of the thing we are trying to decode into, which is why we declare an empty slice of Player the line before.

// Parsing JSON can fail so Decode can return an error.
// There's no point continuing the test if that fails so we check for the error and stop the test with t.Fatalf if it happens.
// Notice that we print the response body along with the error as it's important for someone running the test to see what string cannot be parsed.

// Our endpoint currently does not return a body so it cannot be parsed into JSON.

// Next, we'll want to extend our test so that we can control exactly what data we want back.

// ----------

// We can update the test to assert that the league table contains some players that we will stub in our store.

// Update StubPlayerSstore to let it store a league, which is just a slice of Player. We'll store our expected data in there.

// server_test.go
type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

// Next, update our current test by putting some players in the league property of our stub and assert they get returned from our server.

// You'll need to update the other tests as we have a new field in StubPlayerStore; set it to nil for the other tests.

// func TestLeague(t *testing.T) {
// 	t.Run("it returns the league table as JSON", func(t *testing.T) {
// 		wantedLeague := []Player{
// 			{"Cleo", 32},
// 			{"Chris", 20},
// 			{"Tiest", 14},
// 		}

// 		store := StubPlayerStore{nil, nil, wantedLeague}
// 		server := NewPlayerServer(&store)

// 		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
// 		response := httptest.NewRecorder()

// 		server.ServeHTTP(response, request)

// 		var got []Player

// 		err := json.NewDecoder(response.Body).Decode(&got)

// 		if err != nil {
// 			t.Fatalf("Unable to parse response from server %q into slice of Player, %v", response.Body, err)
// 		}

// 		assertStatus(t, response.Code, http.StatusOK)

// 		if !reflect.DeepEqual(got, wantedLeague) {
// 			t.Errorf("got %v want %v", got, wantedLeague)
// 		}
// 	})
// }

// For StubPlayerStore it's pretty easy, just return the league field we added earlier.

func (s *StubPlayerStore) GetLeague() []Player {
	return s.league
}

// Whilst it would be pretty straightforward to implement GetLeague "properly" by iterating over the map remember we are just trying to write the minimal amount of code to make the tests pass.

// -----------

// The test code does not convey our intent very well and has a lot of boilerplate we can refactor away.

// func TestLeague(t *testing.T) {
// 	t.Run("it returns the league table as JSON", func(t *testing.T) {
// 		wantedLeague := []Player{
// 			{"Cleo", 32},
// 			{"Chris", 20},
// 			{"Tiest", 14},
// 		}

// 		store := StubPlayerStore{nil, nil, wantedLeague}
// 		server := NewPlayerServer(&store)

//		request := newLeagueRequest()
// 		response := httptest.NewRecorder()

// 		server.ServeHTTP(response, request)

// 		got := getLeagueFromResponse(t, response.Body)
// 		assertStatus(t, response.Code, http.StatusOK)
// 		assertLeague(t, got, wantedLeague)
// 	})
// }

// Here are the new helpers

func getLeagueFromResponse(t testing.TB, body io.Reader) (league []Player) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player %v", body, err)
	}

	return
}

func assertLeague(t testing.TB, got, want []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

// One final thing we need to do for our server to work is make sure we return a content-type header in the response so machines can recognise we are returning JSON.

// func TestLeague(t *testing.T) {
// 	t.Run("it returns the league table as JSON", func(t *testing.T) {
// 		wantedLeague := []Player{
// 			{"Cleo", 32},
// 			{"Chris", 20},
// 			{"Tiest", 14},
// 		}

// 		store := StubPlayerStore{nil, nil, wantedLeague}
// 		server := NewPlayerServer(&store)

// 		request := newLeagueRequest()
// 		response := httptest.NewRecorder()

// 		server.ServeHTTP(response, request)

// 		if response.Result().Header.Get("content-type") != "application/json" {
// 			t.Errorf("response did not have content-type of application/json, got %v", response.Result().Header)
// 		}

// 		got := getLeagueFromResponse(t, response.Body)
// 		assertStatus(t, response.Code, http.StatusOK)
// 		assertLeague(t, got, wantedLeague)
// 	})
// }

// Then add a helper for assertContentType.

func assertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.Result().Header)
	}
}

// Use it in the test

func TestLeague(t *testing.T) {
	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := StubPlayerStore{nil, nil, wantedLeague}
		server := NewPlayerServer(&store)

		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Result().Header.Get("content-type") != "application/json" {
			t.Errorf("response did not have content-type of application/json, got %v", response.Result().Header)
		}

		assertStatus(t, response.Code, http.StatusOK)
		assertContentType(t, response, jsonContentType)
		got := getLeagueFromResponse(t, response.Body)
		assertLeague(t, got, wantedLeague)
	})
}
