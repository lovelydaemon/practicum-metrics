package metrics

type MetricsRepo struct{}

func NewMetricsRepo() *MetricsRepo {
	return &MetricsRepo{}
}

func (m *MetricsRepo) Save() {}
