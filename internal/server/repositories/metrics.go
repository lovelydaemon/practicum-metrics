package repositories

import (
	"errors"

	"github.com/lovelydaemon/practicum-metrics/internal/server/storage"
)

type metricsRepo struct {
	storage storage.Storage
}

var (
	ErrRepoNotFound = errors.New("not found")
)

func NewMetricsRepo(storage storage.Storage) *metricsRepo {
	return &metricsRepo{
		storage: storage,
	}
}

func (repo *metricsRepo) GetAll() map[string]any {
	metrics := repo.storage.GetAll()
	return metrics
}

func (repo *metricsRepo) GetGauge(metricName string) (float64, error) {
	value, ok := repo.storage.GetGauge(metricName)
	if !ok {
		return 0, ErrRepoNotFound
	}
	return value, nil
}

func (repo *metricsRepo) GetCounter(metricName string) (int64, error) {
	value, ok := repo.storage.GetCounter(metricName)
	if !ok {
		return 0, ErrRepoNotFound
	}
	return value, nil
}

func (repo *metricsRepo) UpdateGauge(metricName string, metricValue float64) {
	repo.storage.SaveGauge(metricName, metricValue)
}

func (repo *metricsRepo) UpdateCounter(metricName string, metricValue int64) {
	value, ok := repo.storage.GetCounter(metricName)
	if !ok {
		repo.storage.SaveCounter(metricName, metricValue)
		return
	}

	repo.storage.SaveCounter(metricName, metricValue+value)
	return
}
