package storage

type Storage interface {
	SaveGauge(key string, value float64) error
	SaveCounter(key string, value int64) error
	GetGauge(key string) (float64, bool)
	GetCounter(key string) (int64, bool)
	GetAll() map[string]any
}
