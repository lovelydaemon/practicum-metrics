package services

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	MetricTypeGauge   string = "gauge"
	MetricTypeCounter string = "counter"
)

var ErrMetricsUnknownType = errors.New("unknown metric type")
var ErrMetricsEmptyName = errors.New("metric name is empty")
var ErrMetricsNotFound = errors.New("metric not found")

type MetricsService struct {
	repo MetricsRepo
}

func NewMetricsService(repo MetricsRepo) *MetricsService {
	return &MetricsService{
		repo: repo,
	}
}

func (service *MetricsService) GetAll() map[string]any {
	metrics := service.repo.GetAll()
	return metrics
}

func (service *MetricsService) GetMetricValue(metricType, metricName string) (string, error) {
	switch metricType {
	case MetricTypeGauge:
		value, err := service.repo.GetGauge(metricName)
		return fmt.Sprintf("%.3f", value), err
	case MetricTypeCounter:
		value, err := service.repo.GetCounter(metricName)
		return fmt.Sprintf("%d", value), err
	default:
		return "", ErrMetricsUnknownType
	}
}

func (service *MetricsService) Save(metricType, metricName, metricValue string) error {
	name := strings.TrimSpace(metricName)
	if name == "" {
		return ErrMetricsEmptyName
	}

	switch metricType {
	case MetricTypeGauge:
		if err := service.saveGauge(name, metricValue); err != nil {
			return err
		}
	case MetricTypeCounter:
		if err := service.saveCounter(name, metricValue); err != nil {
			return err
		}
	default:
		return ErrMetricsUnknownType
	}

	return nil
}

func (service *MetricsService) saveGauge(metricName, metricValue string) error {
	parsedValue, err := strconv.ParseFloat(metricValue, 64)
	if err != nil {
		return err
	}

	service.repo.SaveGauge(metricName, parsedValue)

	return nil
}

func (service *MetricsService) saveCounter(metricName, metricValue string) error {
	parsedValue, err := strconv.ParseInt(metricValue, 10, 64)
	if err != nil {
		return err
	}

	service.repo.SaveCounter(metricName, parsedValue)

	return nil
}
