package usecase

//import "github.com/lovelydaemon/practicum-metrics/internal/server/storage"

// тут идет конкретная реализация usecase метрик

type Storage interface {
	Save(key string, value any)
	Get(key string) (any, bool)
}

// это что-то не то

func UpdateGaugeMetrics() {

	// тут какая-то логика и хранение

}

func UpdateCounterMetrics() {

	// тут какая-то логика и хранение

}

// сделаем по другому, сначала просто разделим на слои как привыкли и прокинем все необходимое
