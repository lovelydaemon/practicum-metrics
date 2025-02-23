package storage

import "errors"

var ErrStorageEmptyKey = errors.New("Key can't be empty")

type MemStorage struct {
	gauge   map[string]float64
	counter map[string]int64
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		gauge:   make(map[string]float64),
		counter: make(map[string]int64),
	}
}

func (m *MemStorage) GetAll() map[string]any {
	resultSize := len(m.gauge) + len(m.counter)
	result := make(map[string]any, resultSize)

	for k, v := range m.gauge {
		result[k] = v
	}

	for k, v := range m.counter {
		result[k] = v
	}

	return result
}

func (m *MemStorage) SaveGauge(key string, value float64) error {
	if key == "" {
		return ErrStorageEmptyKey
	}

	m.gauge[key] = value

	return nil
}

func (m *MemStorage) GetGauge(key string) (float64, bool) {
	value, ok := m.gauge[key]
	return value, ok
}

func (m *MemStorage) SaveCounter(key string, value int64) error {
	if key == "" {
		return ErrStorageEmptyKey
	}

	m.counter[key] = value

	return nil
}

func (m *MemStorage) GetCounter(key string) (int64, bool) {
	value, ok := m.counter[key]
	return value, ok
}
