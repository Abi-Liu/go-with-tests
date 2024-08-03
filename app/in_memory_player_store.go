package main

import "sync"

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		store: map[string]int{},
		mu:    sync.RWMutex{},
	}
}

type InMemoryPlayerStore struct {
	store map[string]int
	mu    sync.RWMutex
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) GetLeague() League {
	var league []Player

	for name, wins := range i.store {
		league = append(league, Player{Name: name, Wins: wins})
	}

	return league
}
