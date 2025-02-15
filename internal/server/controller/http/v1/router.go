package v1

import (
	"net/http"

	"github.com/lovelydaemon/practicum-metrics/internal/server/services"
)

func NewRouter(mux *http.ServeMux, metrics services.Metrics) http.Handler {
	newMetricsRoutes(mux, metrics)
	return mux
}
