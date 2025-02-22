package services

type (
	Metrics interface {
		UpdateMetrics(metricType, metricName, metricValue string) error
		GetMetricValue(metricType, metricName string) (string, error)
		GetAll() map[string]any
	}

	MetricsRepo interface {
		GetAll() map[string]any
		GetGauge(metricName string) (float64, error)
		GetCounter(metricName string) (int64, error)
		UpdateGauge(metricName string, metricValue float64)
		UpdateCounter(metricName string, metricValue int64)
	}
)
