package main

// I am going to take some liberties here and write more code than you may be comfortable with without writing a test.

// This is allowed!
// We still have a test checking things should be working correctly but it is not around the specific unit we're working with (InMemoryPlayerStore).

// If I were to get stuck in this scenario, I would revert my changes back to the failing test and then write more specific unit tests around InMemoryPlayerStore to help me drive out a solution.

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct {
	store map[string]int
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

// We need to store the data so I've added a map[string]int to the InMemoryPlayerStore struct

// For convenience I've made NewInMemoryPlayerStore to initialise the store, and updated the integration test to use it:
