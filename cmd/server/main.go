package main

import (
	"net/http"
	"strconv"
)

const (
	gauge   = "gauge"
	counter = "counter"
)

// новое значение замещает предыдущее
//type gauge float64

// новое значение добавляется к предыдущему, если какое-то значение уже было
// известно серверу
//type counter int64

type MemStorage struct {
	storage map[string]any
}

func newMemStorage() *MemStorage {
	return &MemStorage{
		storage: make(map[string]any),
	}
}

// TODO посмотреть как правильно ее создавать
//func newMemStorage() *MemStorage {
//	return &MemStorage{}
//}

// описать интерфейс для взаимодействия с хранилищем

var storage = newMemStorage()

func main() {
	//storage := newMemStorage()
	mux := http.NewServeMux()
	mux.HandleFunc("POST /update/{metricType}/{metricName}/{metricValue}", updateMetrics)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}

// Реализуем для хранилища методы для работы с ним

// много повторов
// работа с хранилищем организована плохо, потому что нет отдельных методов

// TODO переделать подключение хендлером, нужно прокинуть в хендлер хранилище

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

	// тут нужен вариант как в js 'gauge' | 'counter'
	switch metricType {
	case gauge:
		if err := updateGauge(metricName, metricValue); err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
	case counter:
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

	storage.storage[metricName] = newValue
	return nil
}

func updateCounter(metricName string, metricValue string) error {
	newValue, err := strconv.ParseInt(metricValue, 10, 64)
	if err != nil {
		return err
	}

	if value, exists := storage.storage[metricName]; exists {
		switch value.(type) {
		case int64:
			intValue := value.(int64)
			storage.storage[metricName] = newValue + intValue
		case float64:
			floatValue := value.(float64)
			storage.storage[metricName] = newValue + int64(floatValue)
		default:

		}
	} else {
		storage.storage[metricName] = newValue
	}

	return nil
}
