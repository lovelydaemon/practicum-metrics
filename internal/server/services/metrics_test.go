package services

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) UpdateGauge(metricName string, metricValue float64) {
	return
}
func (m *MockRepo) UpdateCounter(metricName string, metricValue int64) {
	return
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
