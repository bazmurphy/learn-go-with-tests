package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Now that we have sorted out PlayerServer for now we can turn our attention to InMemoryPlayerStore because right now if we tried to demo this to the product owner /league will not work.

// The quickest way for us to get some confidence is to add to our integration test, we can hit the new endpoint and check we get back the correct response from /league.

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())
		assertStatus(t, response.Code, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)
		want := []Player{
			{"Pepper", 3},
		}
		assertLeague(t, got, want)
	})
}
