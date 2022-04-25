package storage

import (
	"github.com/Farengier/myofficestorage/src/storage/memory"
)

type Storage interface {
	Store(key string, value []byte) error
	Get(key string) ([]byte, error)
	Delete(key string) error
	Close()
}

func MemoryStorage() Storage {
	return memory.New()
}
