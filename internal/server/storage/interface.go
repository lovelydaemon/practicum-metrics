package storage

type Storage interface {
	Save(key string, value any) error
	Get(key string) (any, bool)
}
