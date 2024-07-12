package main

import (
	"log"
	"net/http"
)

// func main() {
// 	server := &PlayerServer{NewInMemoryPlayerStore()}
// 	log.Fatal(http.ListenAndServe(":5000", server))
// }

// Our product owner has a new requirement; to have a new endpoint called /league which returns a list of all players stored.
// She would like this to be returned as JSON.

func main() {
	store := NewInMemoryPlayerStore()
	server := NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":5000", server))
}
