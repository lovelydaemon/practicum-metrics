package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lovelydaemon/practicum-metrics/internal/server/services"
)

func NewRouter(mux *chi.Mux, metrics services.Metrics) http.Handler {
	newMetricsRoutes(mux, metrics)
	return mux
}
