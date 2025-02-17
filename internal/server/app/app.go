package app

import (
	"net/http"

	v1 "github.com/lovelydaemon/practicum-metrics/internal/server/controller/http/v1"
	"github.com/lovelydaemon/practicum-metrics/internal/server/repositories"
	"github.com/lovelydaemon/practicum-metrics/internal/server/services"
	"github.com/lovelydaemon/practicum-metrics/internal/server/storage"
)

func Run() error {
	storage := storage.NewMemStorage()

	mux := http.NewServeMux()

	metricsRepo := repositories.NewMetricsRepo(storage)
	metricsService := services.NewMetricsService(metricsRepo)

	handler := v1.NewRouter(mux, metricsService)

	return http.ListenAndServe(":8080", handler)
}
