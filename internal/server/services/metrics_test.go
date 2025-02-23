package services_test

import (
	"testing"

	"github.com/lovelydaemon/practicum-metrics/internal/server/repositories"
	"github.com/lovelydaemon/practicum-metrics/internal/server/services"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func metrics(t *testing.T) (*services.MetricsService, *MockMetricsRepo) {
	t.Helper()

	ctrl := gomock.NewController(t)
	repo := NewMockMetricsRepo(ctrl)
	metrics := services.NewMetricsService(repo)

	return metrics, repo
}

func TestGetMetricValue(t *testing.T) {
	metrics, repo := metrics(t)

	t.Run("should return gauge value", func(t *testing.T) {
		repo.EXPECT().GetGauge("testGauge").Return(1.111, nil)
		value, _ := metrics.GetMetricValue(services.MetricTypeGauge, "testGauge")
		assert.Equal(t, "1.111", value)
	})

	t.Run("should return counter value", func(t *testing.T) {
		repo.EXPECT().GetCounter("testCounter").Return(int64(10), nil)
		value, _ := metrics.GetMetricValue(services.MetricTypeCounter, "testCounter")
		assert.Equal(t, "10", value)
	})

	t.Run("should return error on unknown metric type", func(t *testing.T) {
		_, err := metrics.GetMetricValue("unknown", "testName")
		assert.ErrorIs(t, err, services.ErrMetricsUnknownType)
	})

	t.Run("should return error 'not found' from repo", func(t *testing.T) {
		repo.EXPECT().GetGauge("testName").Return(0.0, repositories.ErrRepoNotFound)
		_, err := metrics.GetMetricValue(services.MetricTypeGauge, "testName")
		assert.ErrorIs(t, err, repositories.ErrRepoNotFound)
	})
}

func TestSave(t *testing.T) {
	metrics, repo := metrics(t)

	t.Run("should return error on empty metric name", func(t *testing.T) {
		err := metrics.Save(services.MetricTypeGauge, "", "1.111")
		assert.ErrorIs(t, err, services.ErrMetricsEmptyName)
	})

	t.Run("should return error on unknown metric type", func(t *testing.T) {
		err := metrics.Save("unknown", "testName", "1.111")
		assert.ErrorIs(t, err, services.ErrMetricsUnknownType)
	})

	t.Run("should return error when trying to parse gauge value in float64 type", func(t *testing.T) {
		err := metrics.Save(services.MetricTypeGauge, "testGauge", "text")
		assert.Error(t, err)
	})

	t.Run("should return error when trying to parse counter value in int64 type", func(t *testing.T) {
		err := metrics.Save(services.MetricTypeCounter, "testCounter", "text")
		assert.Error(t, err)
	})

	t.Run("should successfully save gauge value", func(t *testing.T) {
		repo.EXPECT().SaveGauge(gomock.Any(), gomock.Any())
		err := metrics.Save(services.MetricTypeGauge, "testGauge", "1.111")
		assert.NoError(t, err)
	})

	t.Run("should successfully save counter value", func(t *testing.T) {
		repo.EXPECT().SaveCounter(gomock.Any(), gomock.Any())
		err := metrics.Save(services.MetricTypeCounter, "testCounter", "10")
		assert.NoError(t, err)
	})
}

func TestGetAll(t *testing.T) {
	metrics, repo := metrics(t)

	want := map[string]any{
		"testCounter": int64(1),
		"testGauge":   float64(1.111),
	}

	t.Run("should return all data from storage", func(t *testing.T) {
		repo.EXPECT().GetAll().Return(want)
		result := metrics.GetAll()
		if !assert.ObjectsAreEqual(want, result) {
			t.Error("Maps are not equal")
		}
	})
}
