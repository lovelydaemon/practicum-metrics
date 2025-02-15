package repositories

import "github.com/lovelydaemon/practicum-metrics/internal/server/storage"

type metricsRepo struct {
	storage storage.Storage
}

func NewMetricsRepo(storage storage.Storage) *metricsRepo {
	return &metricsRepo{
		storage: storage,
	}
}

func (repo *metricsRepo) UpdateGauge(metricName string, metricValue float64) {
	repo.storage.Save(metricName, metricValue)
}

func (repo *metricsRepo) UpdateCounter(metricName string, metricValue int64) {
	if value, exists := repo.storage.Get(metricName); exists {
		switch value.(type) {
		case int64:
			intValue := value.(int64)
			repo.storage.Save(metricName, metricValue+intValue)
		case float64:
			floatValue := value.(float64)
			repo.storage.Save(metricName, metricValue+int64(floatValue))
		default:

		}
	} else {
		repo.storage.Save(metricName, metricValue)
	}
}
