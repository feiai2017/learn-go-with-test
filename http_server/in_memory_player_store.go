package http_server

type InMemoryPlayerStore struct {
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	score := i.store[name]
	return score
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name] ++
}
