package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Integration tests can be useful for testing that larger areas of your system work but you must bear in mind:

// They are harder to write

// When they fail, it can be difficult to know why (usually it's a bug within a component of the integration test) and so can be harder to fix

// They are sometimes slower to run (as they often are used with "real" components, like a database)

// For that reason, it is recommended that you research The Test Pyramid.

// func TestRecordingWinsAndRetrievingThem(t *testing.T) {
// 	store := InMemoryPlayerStore{}
// 	server := PlayerServer{&store}
// 	player := "Pepper"

// 	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
// 	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
// 	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

// 	response := httptest.NewRecorder()
// 	server.ServeHTTP(response, newGetScoreRequest(player))

// 	assertStatus(t, response.Code, http.StatusOK)

// 	assertResponseBody(t, response.Body.String(), "3")
// }

// We are creating our two components we are trying to integrate with: InMemoryPlayerStore and PlayerServer.

// We then fire off 3 requests to record 3 wins for player.
// We're not too concerned about the status codes in this test as it's not relevant to whether they are integrating well.

// The next response we do care about (so we store a variable response) because we are going to try and get the player's score.

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := PlayerServer{store}
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))

	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "3")
}
