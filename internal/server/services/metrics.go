package services

import (
	"errors"
	"strconv"
	"strings"
)

const (
	typeGauge   string = "gauge"
	typeCounter string = "counter"
)

var ErrMetricsUnknownType = errors.New("unknown metric type")
var ErrMetricsEmptyName = errors.New("metric name is empty")

type metricsService struct {
	repo MetricsRepo
}

func NewMetricsService(repo MetricsRepo) *metricsService {
	return &metricsService{
		repo: repo,
	}
}

func (service *metricsService) UpdateMetrics(metricType, metricName, metricValue string) error {
	name := strings.TrimSpace(metricName)
	if name == "" {
		return ErrMetricsEmptyName
	}

	switch metricType {
	case typeGauge:
		if err := service.updateGauge(name, metricValue); err != nil {
			return err
		}
	case typeCounter:
		if err := service.updateCounter(name, metricValue); err != nil {
			return err
		}
	default:
		return ErrMetricsUnknownType
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
