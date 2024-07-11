package main

import (
	"log"
	"net/http"
)

// Create a new main.go file for our application

// func main() {
// 	handler := http.HandlerFunc(PlayerServer)
// 	log.Fatal(http.ListenAndServe(":5000", handler))
// }

// So far all of our application code has been in one file, however, this isn't best practice for larger projects where you'll want to separate things into different files.

// To run this, do go build which will take all the .go files in the directory and build you a program. You can then execute it with ./myprogram.

// http.HandlerFunc

// Earlier we explored that the Handler interface is what we need to implement in order to make a server.
// Typically we do that by creating a struct and make it implement the interface by implementing its own ServeHTTP method.
// However the use-case for structs is for holding data but currently we have no state, so it doesn't feel right to be creating one.

// HandlerFunc lets us avoid this.

// The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers.
// If f is a function with the appropriate signature, HandlerFunc(f) is a Handler that calls f.

// type HandlerFunc func(ResponseWriter, *Request)

// From the documentation, we see that type HandlerFunc has already implemented the ServeHTTP method. By type casting our PlayerServer function with it, we have now implemented the required Handler.

// http.ListenAndServe(":5000"...)

// ListenAndServe takes a port to listen on a Handler.
// If there is a problem the web server will return an error, an example of that might be the port already being listened to.
// For that reason we wrap the call in log.Fatal to log the error to the user.

// What we're going to do now is write another test to force us into making a positive change to try and move away from the hard-coded value.

// ----------

// func main() {
// 	server := &PlayerServer{}
// 	log.Fatal(http.ListenAndServe(":5000", server))
// }

// The program should start up but you'll get a horrible response if you try and hit the server at http://localhost:5000/players/Pepper.

// The reason for this is that we have not passed in a PlayerStore.

// We'll need to make an implementation of one, but that's difficult right now as we're not storing any meaningful data so it'll have to be hard-coded for the time being.

// type InMemoryPlayerStore struct{}

// func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
// 	return 123
// }

// func main() {
// 	server := &PlayerServer{&InMemoryPlayerStore{}}
// 	log.Fatal(http.ListenAndServe(":5000", server))
// }

// -----------

// If you run go build again and hit the same URL you should get "123".
// Not great, but until we store data that's the best we can do.
// It also didn't feel great that our main application was starting up but not actually working.
// We had to manually test to see the problem.

// We have a few options as to what to do next

// - Handle the scenario where the player doesn't exist

// - Handle the POST /players/{name} scenario

// Whilst the POST scenario gets us closer to the "happy path", I feel it'll be easier to tackle the missing player scenario first as we're in that context already.
// We'll get to the rest later.

// -----------

// func (i *InMemoryPlayerStore) RecordWin(name string) {}

// ----------

// The integration test passes, now we just need to change main to use NewInMemoryPlayerStore()

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}

// -----------

// Build it, run it and then use curl to test it out.

// - Run this a few times, change the player names if you like curl -X POST http://localhost:5000/players/Pepper
// - Check scores with curl http://localhost:5000/players/Pepper

// Great! You've made a REST-ish service. To take this forward you'd want to pick a data store to persist the scores longer than the length of time the program runs.

// - Pick a store (Bolt? Mongo? Postgres? File system?)
// - Make PostgresPlayerStore implement PlayerStore
// - TDD the functionality so you're sure it works
// - Plug it into the integration test, check it's still ok
// - Finally plug it into main

// -----------

// We are almost there! Lets take some effort to prevent concurrency errors like these
// fatal error: concurrent map read and map write
// By adding mutexes, we enforce concurrency safety especially for the counter in our RecordWin function.
// Read more about mutexes in the sync chapter.
