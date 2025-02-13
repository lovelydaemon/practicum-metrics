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
	if metricType != gauge && metricType != counter {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	metricName := req.PathValue("metricName")
	if metricName == "" {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	metricValue := req.PathValue("metricValue")

	if metricType == gauge {
		newValue, err := strconv.ParseFloat(metricValue, 64)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		// имя метрики - значение
		// замещение
		storage.storage[metricName] = newValue

		// TODO сохранение в стор
	} else if metricValue == counter {
		newValue, err := strconv.ParseInt(metricValue, 10, 64)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		// имя метрики - значение
		// замещение
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
	}

	res.WriteHeader(http.StatusOK)
}
