package main

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

// So let's just get the compiler happy for now and live with the uncomfortable feeling of an incomplete implementation in our InMemoryStore.

// func (i *InMemoryPlayerStore) GetLeague() []Player {
// 	return nil
// }

// What this is really telling us is that later we're going to want to test this but let's park that for now.

// ----------

// InMemoryPlayerStore is returning nil when you call GetLeague() so we'll need to fix that.

func (i *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return league
}

// All we need to do is iterate over the map and convert each key/value to a Player.
