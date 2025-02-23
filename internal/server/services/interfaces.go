package services

type (
	Metrics interface {
		Save(metricType, metricName, metricValue string) error
		GetMetricValue(metricType, metricName string) (string, error)
		GetAll() map[string]any
	}

	MetricsRepo interface {
		GetAll() map[string]any
		GetGauge(metricName string) (float64, error)
		GetCounter(metricName string) (int64, error)
		SaveGauge(metricName string, metricValue float64)
		SaveCounter(metricName string, metricValue int64)
	}
)
