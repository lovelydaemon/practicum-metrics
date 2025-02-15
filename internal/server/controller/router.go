package controller

import "net/http"

func NewRouter(mux *http.ServeMux) http.Handler {
	// получается тут не база, а общее место для всех вариаций

	// подключаем роуты метрик
	//newMetricsRoutes(mux)

	return mux
}
