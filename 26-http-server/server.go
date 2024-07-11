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

// type PlayerStore interface {
// 	GetPlayerScore(name string) int
// }

// For our PlayerServer to be able to use a PlayerStore, it will need a reference to one.
// Now feels like the right time to change our architecture so that our PlayerServer is now a struct.

type PlayerServer struct {
	store PlayerStore
}

// Finally, we will now implement the Handler interface by adding a method to our new struct and putting in our existing handler code.

// func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	player := strings.TrimPrefix(r.URL.Path, "/players/")
// 	fmt.Fprint(w, p.store.GetPlayerScore(player))
// }

// The only other change is we now call our store.GetPlayerScore to get the score, rather than the local function we defined (which we can now delete).

// This was quite a few changes and we know our tests and application will no longer compile, but just relax and let the compiler work through it.

// We need to change our tests to instead create a new instance of our PlayerServer and then call its method ServeHTTP.

// -----------

// func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	player := strings.TrimPrefix(r.URL.Path, "/players/")

// 	w.WriteHeader(http.StatusNotFound)

// 	fmt.Fprint(w, p.store.GetPlayerScore(player))
// }

// Sometimes I heavily roll my eyes when TDD advocates say "make sure you just write the minimal amount of code to make it pass" as it can feel very pedantic.

// But this scenario illustrates the example well.
// I have done the bare minimum (knowing it is not correct), which is write a StatusNotFound on all responses but all our tests are passing!

// By doing the bare minimum to make the tests pass it can highlight gaps in your tests.
// In our case, we are not asserting that we should be getting a StatusOK when players do exist in the store.

// -------------

// Now our first two tests fail because of the 404 instead of 200, so we can fix PlayerServer to only return not found if the score is 0.

// func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	player := strings.TrimPrefix(r.URL.Path, "/players/")

// 	score := p.store.GetPlayerScore(player)

// 	if score == 0 {
// 		w.WriteHeader(http.StatusNotFound)
// 	}

// 	fmt.Fprint(w, score)
// }

// ------------

// Remember we are deliberately committing sins, so an if statement based on the request's method will do the trick.

// func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPost {
// 		w.WriteHeader(http.StatusAccepted)
// 	}

// 	player := strings.TrimPrefix(r.URL.Path, "/players/")

// 	score := p.store.GetPlayerScore(player)

// 	if score == 0 {
// 		w.WriteHeader(http.StatusNotFound)
// 	}

// 	fmt.Fprint(w, score)
// }

// The handler is looking a bit muddled now. Let's break the code up to make it easier to follow and isolate the different functionality into new functions.

// func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodPost:
// 		p.processWin(w)
// 	case http.MethodGet:
// 		p.showScore(w, r)
// 	}
// }

// func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
// 	player := strings.TrimPrefix(r.URL.Path, "/players/")

// 	score := p.store.GetPlayerScore(player)

// 	if score == 0 {
// 		w.WriteHeader(http.StatusNotFound)
// 	}

// 	fmt.Fprint(w, score)
// }

// func (p *PlayerServer) processWin(w http.ResponseWriter) {
// 	w.WriteHeader(http.StatusAccepted)
// }

// This makes the routing aspect of ServeHTTP a bit clearer and means our next iterations on storing can just be inside processWin.

// Next, we want to check that when we do our POST /players/{name} that our PlayerStore is told to record the win.

// ----------

// We need to update PlayerServer's idea of what a PlayerStore is by changing the interface if we're going to be able to call RecordWin

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

// Now that PlayerStore has RecordWin we can call it within our PlayerServer

// func (p *PlayerServer) processWin(w http.ResponseWriter) {
// 	p.store.RecordWin("Bob")
// 	w.WriteHeader(http.StatusAccepted)
// }

// Run the tests and it should be passing! Obviously "Bob" isn't exactly what we want to send to RecordWin, so let's further refine the test.

// -----------

// func (p *PlayerServer) processWin(w http.ResponseWriter, r *http.Request) {
// 	player := strings.TrimPrefix(r.URL.Path, "/players/")
// 	p.store.RecordWin(player)
// 	w.WriteHeader(http.StatusAccepted)
// }

// We changed processWin to take http.Request so we can look at the URL to extract the player's name.
// Once we have that we can call our store with the correct value to make the test pass.

// -----------

// We can DRY up this code a bit as we're extracting the player name the same way in two places

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

// Even though our tests are passing we don't really have working software.
// If you try and run main and use the software as intended it doesn't work because we haven't got round to implementing PlayerStore correctly.
// This is fine though; by focusing on our handler we have identified the interface that we need, rather than trying to design it up-front.

// We could start writing some tests around our InMemoryPlayerStore but it's only here temporarily until we implement a more robust way of persisting player scores (i.e. a database).

// What we'll do for now is write an integration test between our PlayerServer and InMemoryPlayerStore to finish off the functionality.
// This will let us get to our goal of being confident our application is working, without having to directly test InMemoryPlayerStore.
// Not only that, but when we get around to implementing PlayerStore with a database, we can test that implementation with the same integration test.
