package v1

import (
	"net/http"

	"github.com/lovelydaemon/practicum-metrics/internal/server/services"
)

const (
	MetricTypeGauge   string = "gauge"
	MetricTypeCounter string = "counter"
)

type metricsRoutes struct {
	service services.Metrics
}

func newMetricsRoutes(mux *http.ServeMux, service services.Metrics) {
	c := &metricsRoutes{
		service: service,
	}

	mux.HandleFunc("POST /update/{metricType}/{metricName}/{metricValue}", c.updateMetrics)
}

func (r *metricsRoutes) updateMetrics(res http.ResponseWriter, req *http.Request) {
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

	err := r.service.UpdateMetrics(metricType, metricName, metricValue)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	res.WriteHeader(http.StatusOK)
}
