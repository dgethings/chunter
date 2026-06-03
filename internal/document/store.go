package document

import (
	"fmt"
	"sync"
)

type Store struct {
	mu   sync.RWMutex
	docs map[string]*Document
}

func NewStore() *Store {
	return &Store{
		docs: make(map[string]*Document),
	}
}

func (s *Store) Get(uri string) (*Document, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	doc, ok := s.docs[uri]
	if !ok {
		return nil, fmt.Errorf("document not found: %s", uri)
	}
	return doc, nil
}

func (s *Store) Put(doc *Document) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.docs[doc.URI] = doc
}

func (s *Store) Delete(uri string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.docs, uri)
}
