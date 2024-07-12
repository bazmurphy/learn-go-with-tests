package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}
	if name == "Floyd" {
		return "10"
	}
	return ""
}

// type PlayerServer struct {
// 	store PlayerStore
// }

// type PlayerStore interface {
// 	GetPlayerScore(name string) int
// 	RecordWin(name string)
// }

// func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	player := strings.TrimPrefix(r.URL.Path, "/players/")

// 	switch r.Method {
// 	case http.MethodPost:
// 		p.processWin(w, player)
// 	case http.MethodGet:
// 		p.showScore(w, player)
// 	}
// }

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

// Go has a built-in routing mechanism called ServeMux (request multiplexer) which lets you attach http.Handlers to particular request paths.

// Let's commit some sins and get the tests passing in the quickest way we can, knowing we can refactor it with safety once we know the tests are passing.

// func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	router := http.NewServeMux()

// 	router.Handle("/league", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(http.StatusOK)
// 	}))

// 	router.Handle("/players/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		player := strings.TrimPrefix(r.URL.Path, "/players/")

// 		switch r.Method {
// 		case http.MethodPost:
// 			p.processWin(w, player)
// 		case http.MethodGet:
// 			p.showScore(w, player)
// 		}

// 		router.ServeHTTP(w, r)
// 	}))
// }

// When the request starts we create a router and then we tell it for x path use y handler.

// So for our new endpoint, we use http.HandlerFunc and an anonymous function to w.WriteHeader(http.StatusOK) when /league is requested to make our new test pass.

// For the /players/ route we just cut and paste our code into another http.HandlerFunc.

// Finally, we handle the request that came in by calling our new router's ServeHTTP (notice how ServeMux is also an http.Handler?)

// ----------

// ServeHTTP is looking quite big, we can separate things out a bit by refactoring our handlers into separate methods.

// func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	router := http.NewServeMux()
// 	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
// 	router.Handle("/p;layers/", http.HandlerFunc(p.playersHandler))

// 	router.ServeHTTP(w, r)
// }

// func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// }

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

// It's quite odd (and inefficient) to be setting up a router as a request comes in and then calling it.
// What we ideally want to do is have some kind of NewPlayerServer function which will take our dependencies and do the one-time setup of creating the router.
// Each request can then just use that one instance of the router.

// type PlayerServer struct {
// 	store  PlayerStore
// 	router *http.ServeMux
// }

// func NewPlayerServer(store PlayerStore) *PlayerServer {
// 	p := &PlayerServer{
// 		store,
// 		http.NewServeMux(),
// 	}

// 	p.router.Handle("/league", http.HandlerFunc(p.leagueHandler))
// 	p.router.Handle("/players/", http.HandlerFunc(p.playersHandler))

// 	return p
// }

// func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	p.router.ServeHTTP(w, r)
// }

// PlayerServer now needs to store a router.
// We have moved the routing creation out of ServeHTTP and into our NewPlayerServer so this only has to be done once, not per request.
// You will need to update all the test and production code where we used to do PlayerServer{&store} with NewPlayerServer(&store).

// -----------

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	p.Handler = router

	return p
}

// We changed the second property of PlayerServer, removing the named property router http.ServeMux and replaced it with http.Handler; this is called embedding.

// Go does not provide the typical, type-driven notion of subclassing, but it does have the ability to “borrow” pieces of an implementation by embedding types within a struct or interface.

// What this means is that our PlayerServer now has all the methods that http.Handler has, which is just ServeHTTP.

// To "fill in" the http.Handler we assign it to the router we create in NewPlayerServer.
// We can do this because http.ServeMux has the method ServeHTTP.

// This lets us remove our own ServeHTTP method, as we are already exposing one via the embedded type.

// Embedding is a very interesting language feature. You can use it with interfaces to compose new interfaces.

// type Animal interface {
// 	Eater
// 	Sleeper
// }

// And you can use it with concrete types too, not just interfaces.
// As you'd expect if you embed a concrete type you'll have access to all its public methods and fields.

// You must be careful with embedding types because you will expose all public methods and fields of the type you embed.
// In our case, it is ok because we embedded just the interface that we wanted to expose (http.Handler).

// If we had been lazy and embedded http.ServeMux instead (the concrete type) it would still work but users of PlayerServer would be able to add new routes to our server because Handle(path, handler) would be public.

// When embedding types, really think about what impact that has on your public API.

// It is a very common mistake to misuse embedding and end up polluting your APIs and exposing the internals of your type.

// Now we've restructured our application we can easily add new routes and have the start of the /league endpoint.
// We now need to make it return some useful information.

// We should return some JSON that looks something like this.

// [
//    {
//       "Name":"Bill",
//       "Wins":10
//    },
//    {
//       "Name":"Alice",
//       "Wins":15
//    }
// ]

type Player struct {
	Name string
	Wins int
}

// Our endpoint currently does not return a body so it cannot be parsed into JSON.

// func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
// 	leagueTable := []Player{
// 		{"Chris", 20},
// 	}

// 	json.NewEncoder(w).Encode(leagueTable)

// 	w.WriteHeader(http.StatusOK)
// }

// Encoding and Decoding

// Notice the lovely symmetry in the standard library.

// To create an Encoder you need an io.Writer which is what http.ResponseWriter implements.

// To create a Decoder you need an io.Reader which the Body field of our response spy implements.

// Throughout this book, we have used io.Writer and this is another demonstration of its prevalence in the standard library and how a lot of libraries easily work with it.

// ----------

// It would be nice to introduce a separation of concern between our handler and getting the leagueTable as we know we're going to not hard-code that very soon.

// func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
// 	json.NewEncoder(w).Encode(p.getLeagueTable())
// 	w.WriteHeader(http.StatusOK)
// }

func (p *PlayerServer) getLeagueTable() []Player {
	return []Player{
		{"Chris", 20},
	}
}

// Next, we'll want to extend our test so that we can control exactly what data we want back.

// We know the data is in our StubPlayerStore and we've abstracted that away into an interface PlayerStore.
// We need to update this so anyone passing us in a PlayerStore can provide us with the data for leagues.

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []Player
}

// Now we can update our handler code to call that rather than returning a hard-coded list.
// Delete our method getLeagueTable() and then update leagueHandler to call GetLeague().

// func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
// 	json.NewEncoder(w).Encode(p.store.GetLeague())
// 	w.WriteHeader(http.StatusOK)
// }

// The compiler is complaining because InMemoryPlayerStore and StubPlayerStore do not have the new method we added to our interface.

// For StubPlayerStore it's pretty easy, just return the league field we added earlier.

// ----------

// One final thing we need to do for our server to work is make sure we return a content-type header in the response so machines can recognise we are returning JSON.

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(p.store.GetLeague())
}

// Create a constant for "application/json" and use it in leagueHandler

const jsonContentType = "application/json"
