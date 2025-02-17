package v1

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lovelydaemon/practicum-metrics/internal/server/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) UpdateMetrics(metricType, metricName, metricValue string) error {
	args := m.Called(metricType, metricName, metricValue)
	return args.Error(0)
}

func TestMetrics_updateMetrics(t *testing.T) {
	handler := http.NewServeMux()

	mockService := new(MockService)

	mockService.On("UpdateMetrics", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	newMetricsRoutes(handler, mockService)

	srv := httptest.NewServer(handler)
	defer srv.Close()

	client := srv.Client()

	type request struct {
		method      string
		contentType string
	}

	tests := []struct {
		name         string
		request      request
		expectedCode int
	}{
		{
			name: "GET method not allowed",
			request: request{
				method: http.MethodGet,
			},
			expectedCode: http.StatusMethodNotAllowed,
		},
		{
			name: "PATCH method not allowed",
			request: request{
				method: http.MethodPatch,
			},
			expectedCode: http.StatusMethodNotAllowed,
		},
		{
			name: "PUT method not allowed",
			request: request{
				method: http.MethodPut,
			},
			expectedCode: http.StatusMethodNotAllowed,
		},
		{
			name: "DELETE method not allowed",
			request: request{
				method: http.MethodDelete,
			},
			expectedCode: http.StatusMethodNotAllowed,
		},
		{
			name: "bad content type",
			request: request{
				method:      http.MethodPost,
				contentType: "application/json",
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "successful case",
			request: request{
				method:      http.MethodPost,
				contentType: "text/plain",
			},
			expectedCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.request.method, srv.URL+"/update/gauge/name/123", nil)
			require.NoError(t, err)

			req.Header.Set("content-type", tt.request.contentType)

			res, err := client.Do(req)
			require.NoError(t, err)

			_, err = io.ReadAll(res.Body)
			require.NoError(t, err)

			defer res.Body.Close()

			assert.Equal(t, tt.expectedCode, res.StatusCode)
		})
	}
}

func TestMetrics_updateMetrics_notFound(t *testing.T) {
	handler := http.NewServeMux()

	mockService := new(MockService)

	mockService.On("UpdateMetrics", mock.Anything, mock.Anything, mock.Anything).Return(services.ErrMetricsEmptyName)

	newMetricsRoutes(handler, mockService)

	srv := httptest.NewServer(handler)
	defer srv.Close()

	client := srv.Client()

	t.Run("metric name is empty", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, srv.URL+"/update/gauge/name/123", nil)
		require.NoError(t, err)

		req.Header.Set("content-type", "text/plain")

		res, err := client.Do(req)
		require.NoError(t, err)

		_, err = io.ReadAll(res.Body)
		require.NoError(t, err)

		defer res.Body.Close()

		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
}

func TestMetrics_updateMetrics_badRequest(t *testing.T) {
	handler := http.NewServeMux()

	mockService := new(MockService)

	mockService.On("UpdateMetrics", mock.Anything, mock.Anything, mock.Anything).Return(services.ErrMetricsUnknownType)

	newMetricsRoutes(handler, mockService)

	srv := httptest.NewServer(handler)
	defer srv.Close()

	client := srv.Client()

	t.Run("metric type is unknown", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, srv.URL+"/update/gauge/name/123", nil)
		require.NoError(t, err)

		req.Header.Set("content-type", "text/plain")

		res, err := client.Do(req)
		require.NoError(t, err)

		_, err = io.ReadAll(res.Body)
		require.NoError(t, err)

		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})
}
