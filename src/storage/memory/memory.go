package memory

import (
	"errors"
	"sync"
)

type ms struct {
	sync.RWMutex
	data map[string][]byte
}

func New() *ms {
	return &ms{
		data: make(map[string][]byte),
	}
}

func (s *ms) Store(key string, value []byte) error {
	s.Lock()
	defer s.Unlock()

	s.data[key] = value
	return nil
}
func (s *ms) Get(key string) ([]byte, error) {
	s.RLock()
	defer s.RUnlock()

	v, ok := s.data[key]
	if !ok {
		return nil, errors.New("key not found")
	}
	return v, nil
}
func (s *ms) Delete(key string) error {
	s.Lock()
	defer s.Unlock()

	delete(s.data, key)
	return nil
}
func (s *ms) Close() {
}
