package services

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/lovelydaemon/practicum-metrics/internal/server/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) GetGauge(metricName string) (float64, error) {
	args := m.Called(metricName)
	return float64(args.Get(0).(float64)), args.Error(1)
}

func (m *MockRepo) GetCounter(metricName string) (int64, error) {
	args := m.Called(metricName)
	return int64(args.Int(0)), args.Error(1)
}

func (m *MockRepo) UpdateGauge(metricName string, metricValue float64) {
	return
}

func (m *MockRepo) UpdateCounter(metricName string, metricValue int64) {
	return
}

func (m *MockRepo) GetAll() map[string]any {
	args := m.Called()
	return args.Get(0).(map[string]any)
}

func TestGetMetricValue_GetValueSuccess(t *testing.T) {
	mockRepo := new(MockRepo)
	service := NewMetricsService(mockRepo)

	mockRepo.On("GetGauge", "test").Return(1.123, nil)
	mockRepo.On("GetCounter", "test").Return(123, nil)

	tests := []struct {
		name       string
		metricType string
		want       string
	}{
		{
			name:       "return gauge value",
			metricType: metricTypeGauge,
			want:       "1.123",
		},
		{
			name:       "return counter value",
			metricType: metricTypeCounter,
			want:       "123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, _ := service.GetMetricValue(tt.metricType, "test")
			assert.Equal(t, tt.want, value)
		})
	}
}

func TestGetMetricValue_GetValueFailed(t *testing.T) {
	mockRepo := new(MockRepo)
	service := NewMetricsService(mockRepo)

	mockRepo.On("GetGauge", "test").Return(0.0, repositories.ErrRepoNotFound)
	mockRepo.On("GetCounter", "test").Return(0, repositories.ErrRepoNotFound)

	tests := []struct {
		name       string
		metricType string
		error      error
	}{
		{
			name:       "unknown metric type",
			metricType: "test",
			error:      ErrMetricsUnknownType,
		},
		{
			name:       "gauge metric not found",
			metricType: metricTypeGauge,
			error:      repositories.ErrRepoNotFound,
		},
		{
			name:       "counter metric not found",
			metricType: metricTypeCounter,
			error:      repositories.ErrRepoNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := service.GetMetricValue(tt.metricType, "test")
			assert.ErrorIs(t, err, tt.error)
		})
	}
}

func TestUpdateMetrics(t *testing.T) {
	mockRepo := new(MockRepo)
	service := NewMetricsService(mockRepo)

	type args struct {
		metricType  string
		metricName  string
		metricValue string
	}

	tests := []struct {
		name      string
		args      args
		wantError bool
		error     error
	}{
		{
			name: "empty name",
			args: args{
				metricType:  "",
				metricName:  "",
				metricValue: "",
			},
			wantError: true,
			error:     ErrMetricsEmptyName,
		},
		{
			name: "unknown type",
			args: args{
				metricType:  "type",
				metricName:  "name",
				metricValue: "123",
			},
			wantError: true,
			error:     ErrMetricsUnknownType,
		},
		{
			name: "parse to float64 error",
			args: args{
				metricType:  "gauge",
				metricName:  "name",
				metricValue: "test",
			},
			wantError: true,
			error:     strconv.ErrSyntax,
		},
		{
			name: "parse to int64 error",
			args: args{
				metricType:  "counter",
				metricName:  "name",
				metricValue: "test",
			},
			wantError: true,
			error:     strconv.ErrSyntax,
		},
		{
			name: "successful case for gauge type",
			args: args{
				metricType:  "gauge",
				metricName:  "name",
				metricValue: "1.1",
			},
			wantError: false,
		},
		{
			name: "successful case for counter type",
			args: args{
				metricType:  "counter",
				metricName:  "name",
				metricValue: "1",
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.UpdateMetrics(tt.args.metricType, tt.args.metricName, tt.args.metricValue)

			if tt.wantError {
				assert.ErrorIs(t, err, tt.error)
				return
			}

			assert.NoError(t, err)
		})
	}
}

func TestGetAll(t *testing.T) {
	mockRepo := new(MockRepo)
	service := NewMetricsService(mockRepo)

	want := map[string]any{
		"testCounter": int64(1),
		"testGauge":   float64(1.1),
	}

	mockRepo.On("GetAll").Return(want)

	t.Run("should return all data from storage", func(t *testing.T) {
		result := service.GetAll()
		if !reflect.DeepEqual(result, want) {
			t.Error("Maps are not equal")
		}
	})
}
