package services

type (
	Metrics interface {
		UpdateMetrics(metricType, metricName, metricValue string) error
	}

	MetricsRepo interface {
		UpdateGauge(metricName string, metricValue float64)
		UpdateCounter(metricName string, metricValue int64)
	}
)
