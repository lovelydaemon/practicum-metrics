package storage

import "errors"

var ErrStorageEmptyKey = errors.New("Key can't be empty")

type MemStorage struct {
	storage map[string]any
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		storage: make(map[string]any),
	}
}

func (m *MemStorage) Save(key string, value any) error {
	if key == "" {
		return ErrStorageEmptyKey
	}

	m.storage[key] = value

	return nil
}

func (m *MemStorage) Get(key string) (any, bool) {
	value, ok := m.storage[key]
	return value, ok
}
