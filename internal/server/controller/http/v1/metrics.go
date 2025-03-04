package v1

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lovelydaemon/practicum-metrics/internal/server/services"
)

type metricsRoutes struct {
	service services.Metrics
}

func newMetricsRoutes(mux *chi.Mux, service services.Metrics) {
	r := &metricsRoutes{
		service: service,
	}

	mux.Get("/", r.getMetricsHTML)
	mux.Get("/value/{metricType}/{metricName}", r.getMetricValue)
	mux.Post("/update/{metricType}/{metricName}/{metricValue}", r.updateMetrics)
}

func (r *metricsRoutes) getMetricsHTML(res http.ResponseWriter, req *http.Request) {
	metrics := r.service.GetAll()

	tmpl, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(res, metrics); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (r *metricsRoutes) getMetricValue(res http.ResponseWriter, req *http.Request) {
	metricType := chi.URLParam(req, "metricType")
	metricName := chi.URLParam(req, "metricName")

	metricValue, err := r.service.GetMetricValue(metricType, metricName)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	payload := []byte(metricValue)

	res.Header().Set("content-length", fmt.Sprintf("%d", len(payload)))
	res.Header().Set("content-type", "text/plain")
	res.WriteHeader(http.StatusOK)
	res.Write(payload)
}

func (r *metricsRoutes) updateMetrics(res http.ResponseWriter, req *http.Request) {
	// broken autotests
	//	contentType := req.Header.Get("content-type")
	//	if contentType != "text/plain" {
	//		res.WriteHeader(http.StatusBadRequest)
	//		return
	//	}

	metricType := chi.URLParam(req, "metricType")
	metricName := chi.URLParam(req, "metricName")
	metricValue := chi.URLParam(req, "metricValue")

	err := r.service.Save(metricType, metricName, metricValue)
	if err != nil {
		if errors.Is(err, services.ErrMetricsEmptyName) {
			res.WriteHeader(http.StatusNotFound)
			return
		}

		res.WriteHeader(http.StatusBadRequest)
		return
	}

	res.WriteHeader(http.StatusOK)
}
