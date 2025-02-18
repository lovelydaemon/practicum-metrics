package app

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	v1 "github.com/lovelydaemon/practicum-metrics/internal/server/controller/http/v1"
	"github.com/lovelydaemon/practicum-metrics/internal/server/repositories"
	"github.com/lovelydaemon/practicum-metrics/internal/server/services"
	"github.com/lovelydaemon/practicum-metrics/internal/server/storage"
	"github.com/lovelydaemon/practicum-metrics/pkg/httpserver"
)

func Run() {
	storage := storage.NewMemStorage()

	mux := http.NewServeMux()

	metricsRepo := repositories.NewMetricsRepo(storage)
	metricsService := services.NewMetricsService(metricsRepo)

	handler := v1.NewRouter(mux, metricsService)

	srv := httpserver.New(handler)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		close(interrupt)
		fmt.Printf("stopped by %s signal\n", s)
	case err := <-srv.Notify():
		fmt.Println(err)
	}

	err := srv.Shutdown()
	if err != nil {
		fmt.Println(err)
	}
}
