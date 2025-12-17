package storage

import "sync"

type Store struct {
	mu sync.Mutex
}

func New() *Store {
	return &Store{}
}
