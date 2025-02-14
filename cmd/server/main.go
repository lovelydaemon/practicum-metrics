package main

import (
	"net/http"
	"strconv"
)

const (
	MetricTypeGauge   string = "gauge"
	MetricTypeCounter string = "counter"
)

var storage = newMemStorage()

type MemStorage struct {
	storage map[string]any
}

func newMemStorage() *MemStorage {
	return &MemStorage{
		storage: make(map[string]any),
	}
}

func (m *MemStorage) Save(key string, value any) {
	m.storage[key] = value
}

func (m *MemStorage) Get(key string) (any, bool) {
	value, ok := m.storage[key]
	return value, ok
}

type Storage interface {
	Save(key string, value any)
	Get(key string) (any, bool)
}

func main() {
	//storage := newMemStorage()
	mux := http.NewServeMux()
	mux.HandleFunc("POST /update/{metricType}/{metricName}/{metricValue}", updateMetrics)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}

// updateMetrics это контроллер, а нам нужно сделать более раскрытую архитектуру
func updateMetrics(res http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("content-type")
	if contentType != "text/plain" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	metricType := req.PathValue("metricType")
	metricName := req.PathValue("metricName")
	if metricName == "" {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	metricValue := req.PathValue("metricValue")

	switch metricType {
	case MetricTypeGauge:
		if err := updateGauge(metricName, metricValue); err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
	case MetricTypeCounter:
		if err := updateCounter(metricName, metricValue); err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
	default:
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	res.WriteHeader(http.StatusOK)
}

func updateGauge(metricName string, metricValue string) error {
	newValue, err := strconv.ParseFloat(metricValue, 64)
	if err != nil {
		return err
	}

	storage.Save(metricName, newValue)

	return nil
}

func updateCounter(metricName string, metricValue string) error {
	newValue, err := strconv.ParseInt(metricValue, 10, 64)
	if err != nil {
		return err
	}

	if value, exists := storage.Get(metricName); exists {
		switch value.(type) {
		case int64:
			intValue := value.(int64)
			storage.Save(metricName, newValue+intValue)
		case float64:
			floatValue := value.(float64)
			storage.Save(metricName, newValue+int64(floatValue))
		default:

		}
	} else {
		storage.Save(metricName, newValue)
	}

	return nil
}
