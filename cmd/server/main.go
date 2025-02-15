package main

import (
	"net/http"

	//"github.com/lovelydaemon/practicum-metrics/internal/server/controller"
	"github.com/lovelydaemon/practicum-metrics/internal/server/metrics"
	"github.com/lovelydaemon/practicum-metrics/internal/server/storage"
)

func main() {
	// общий стор
	storage := storage.NewMemStorage()

	// движок роутинга
	mux := http.NewServeMux()

	// сервис метрик
	metricsService := metrics.NewMetricsService(storage)
	// роутинг метрик
	metrics.NewMetricsController(mux, metricsService)

	//handler := controller.NewRouter(mux)

	// передали все роуты в сервер
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
