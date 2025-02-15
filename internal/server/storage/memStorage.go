package storage

type MemStorage struct {
	storage map[string]any
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		storage: make(map[string]any),
	}
}

func (m *MemStorage) Save(key string, value any) {
	m.storage[key] = value
}

func (m *MemStorage) Get(key string) (any, bool) {
	value, ok := m.storage[key]
	return value, ok
}
