package services

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	metricTypeGauge   string = "gauge"
	metricTypeCounter string = "counter"
)

var ErrMetricsUnknownType = errors.New("unknown metric type")
var ErrMetricsEmptyName = errors.New("metric name is empty")
var ErrMetricsNotFound = errors.New("metric not found")

type metricsService struct {
	repo MetricsRepo
}

func NewMetricsService(repo MetricsRepo) *metricsService {
	return &metricsService{
		repo: repo,
	}
}

func (service *metricsService) GetAll() map[string]any {
	metrics := service.repo.GetAll()
	return metrics
}

func (service *metricsService) GetMetricValue(metricType, metricName string) (string, error) {
	switch metricType {
	case metricTypeGauge:
		value, err := service.repo.GetGauge(metricName)
		return fmt.Sprintf("%.3f", value), err
	case metricTypeCounter:
		value, err := service.repo.GetCounter(metricName)
		return fmt.Sprintf("%d", value), err
	default:
		return "", ErrMetricsUnknownType
	}
}

func (service *metricsService) UpdateMetrics(metricType, metricName, metricValue string) error {
	name := strings.TrimSpace(metricName)
	if name == "" {
		return ErrMetricsEmptyName
	}

	switch metricType {
	case metricTypeGauge:
		if err := service.updateGauge(name, metricValue); err != nil {
			return err
		}
	case metricTypeCounter:
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
