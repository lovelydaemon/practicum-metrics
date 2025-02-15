package metrics

import (
	"strconv"

	"github.com/lovelydaemon/practicum-metrics/internal/server/storage"
)

type MetricsService struct {
	storage *storage.MemStorage
}

func NewMetricsService(storage *storage.MemStorage) *MetricsService {
	return &MetricsService{
		storage: storage,
	}
}

func (m *MetricsService) UpdateGauge(metricName, metricValue string) error {
	newValue, err := strconv.ParseFloat(metricValue, 64)
	if err != nil {
		return err
	}

	m.storage.Save(metricName, newValue)

	return nil

}
func (m *MetricsService) sasi() {}

func (m *MetricsService) UpdateCounter(metricName, metricValue string) error {
	newValue, err := strconv.ParseInt(metricValue, 10, 64)
	if err != nil {
		return err
	}

	if value, exists := m.storage.Get(metricName); exists {
		switch value.(type) {
		case int64:
			intValue := value.(int64)
			m.storage.Save(metricName, newValue+intValue)
		case float64:
			floatValue := value.(float64)
			m.storage.Save(metricName, newValue+int64(floatValue))
		default:

		}
	} else {
		m.storage.Save(metricName, newValue)
	}

	return nil

}
