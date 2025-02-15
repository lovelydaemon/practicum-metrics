package metrics

import (
	"net/http"
)

const (
	MetricTypeGauge   string = "gauge"
	MetricTypeCounter string = "counter"
)

type MetricsController struct {
	service *MetricsService
}

func NewMetricsController(mux *http.ServeMux, service *MetricsService) {
	c := &MetricsController{
		service: service,
	}

	mux.HandleFunc("POST /update/{metricType}/{metricName}/{metricValue}", c.UpdateMetrics)
}

func (m *MetricsController) UpdateMetrics(res http.ResponseWriter, req *http.Request) {
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
		if err := m.service.UpdateGauge(metricName, metricValue); err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
	case MetricTypeCounter:
		if err := m.service.UpdateCounter(metricName, metricValue); err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
	default:
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	res.WriteHeader(http.StatusOK)
}
