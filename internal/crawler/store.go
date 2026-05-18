package crawler

import "sync"

type URLStore struct {
	visited map[string]bool
	mu      sync.Mutex
}

func NewURLStore() *URLStore {
	return &URLStore{
		visited: make(map[string]bool),
	}
}

func (s *URLStore) Add(url string) bool {

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.visited[url] {
		return false
	}

	s.visited[url] = true
	return true
}
