package main

import (
	"fmt"
	"net/http"
	"strings"
)

// To create a web server in Go you will typically call ListenAndServe.

// func ListenAndServe(addr string, handler Handler) error

// This will start a web server listening on a port, creating a goroutine for every request and running it against a Handler.

// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }

// A type implements the Handler interface by implementing the ServeHTTP method which expects two arguments, the first is where we write our response and the second is the HTTP request that was sent to the server.

// -----------

// Create a file named server.go and define PlayerServer

// func PlayerServer(w http.ResponseWriter, r *http.Request) {

// }

// From the DI chapter, we touched on HTTP servers with a Greet function.
// We learned that net/http's ResponseWriter also implements io Writer so we can use fmt.Fprint to send strings as HTTP responses.

// func PlayerServer(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "20")
// }

// We want to wire this up into an application.
// This is important because:
// We'll have actual working software, we don't want to write tests for the sake of it, it's good to see the code in action.
// As we refactor our code, it's likely we will change the structure of the program.
// We want to make sure this is reflected in our application too as part of the incremental approach.

// -----------

// func PlayerServer(w http.ResponseWriter, r *http.Request) {
// 	player := strings.TrimPrefix(r.URL.Path, "/players/")

// 	if player == "Pepper" {
// 		fmt.Fprint(w, "20")
// 		return
// 	}

// 	if player == "Floyd" {
// 		fmt.Fprint(w, "10")
// 		return
// 	}
// }

// This test has forced us to actually look at the request's URL and make a decision.
// So whilst in our heads, we may have been worrying about player stores and interfaces the next logical step actually seems to be about routing.

// If we had started with the store code the amount of changes we'd have to do would be very large compared to this.
// This is a smaller step towards our final goal and was driven by tests.

// We're resisting the temptation to use any routing libraries right now, just the smallest step to get our test passing.

// r.URL.Path returns the path of the request which we can then use strings.TrimPrefix to trim away /players/ to get the requested player.
// It's not very robust but will do the trick for now.

// We can simplify the PlayerServer by separating out the score retrieval into a function

// func PlayerServer(w http.ResponseWriter, r *http.Request) {
// 	player := strings.TrimPrefix(r.URL.Path, "/players/")

// 	fmt.Fprint(w, GetPlayerScore(player))
// }

func GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}
	if name == "Floyd" {
		return "10"
	}
	return ""
}

// ----------

// However, we still shouldn't be happy. It doesn't feel right that our server knows the scores.

// Our refactoring has made it pretty clear what to do.

// We moved the score calculation out of the main body of our handler into a function GetPlayerScore. This feels like the right place to separate the concerns using interfaces.

// Let's move our function we re-factored to be an interface instead

type PlayerStore interface {
	GetPlayerScore(name string) int
}

// For our PlayerServer to be able to use a PlayerStore, it will need a reference to one.
// Now feels like the right time to change our architecture so that our PlayerServer is now a struct.

type PlayerServer struct {
	store PlayerStore
}

// Finally, we will now implement the Handler interface by adding a method to our new struct and putting in our existing handler code.

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, p.store.GetPlayerScore(player))
}

// The only other change is we now call our store.GetPlayerScore to get the score, rather than the local function we defined (which we can now delete).

// This was quite a few changes and we know our tests and application will no longer compile, but just relax and let the compiler work through it.

// We need to change our tests to instead create a new instance of our PlayerServer and then call its method ServeHTTP.
