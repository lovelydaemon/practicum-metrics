package services

import (
	"strconv"
)

const (
	metricTypeGauge   string = "gauge"
	metricTypeCounter string = "counter"
)

type metricsService struct {
	repo MetricsRepo
}

func NewMetricsService(repo MetricsRepo) *metricsService {
	return &metricsService{
		repo: repo,
	}
}

func (service *metricsService) UpdateMetrics(metricType, metricName, metricValue string) error {
	switch metricType {
	case metricTypeGauge:
		if err := service.updateGauge(metricName, metricValue); err != nil {
			return err
		}
	case metricTypeCounter:
		if err := service.updateCounter(metricName, metricValue); err != nil {
			return err
		}
	default:
		return nil
	}

	return nil
}

func (service *metricsService) updateGauge(metricName, metricValue string) error {
	parsedValue, err := strconv.ParseFloat(metricValue, 64)
	if err != nil {
		return err
	}

	service.repo.UpdateGauge(metricName, parsedValue)

	return nil
}

func (service *metricsService) updateCounter(metricName, metricValue string) error {
	parsedValue, err := strconv.ParseInt(metricValue, 10, 64)
	if err != nil {
		return err
	}

	service.repo.UpdateCounter(metricName, parsedValue)

	return nil
}
