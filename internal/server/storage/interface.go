package storage

type Storage interface {
	Save(key string, value any)
	Get(key string) (any, bool)
}
