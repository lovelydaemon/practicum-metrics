package v1

// Скрыто до лучших времен

// "errors"
// "io"
// "net/http"
// "net/http/httptest"
// "testing"
//
// "github.com/lovelydaemon/practicum-metrics/internal/server/services"
// "github.com/stretchr/testify/assert"
// "github.com/stretchr/testify/mock"
// "github.com/stretchr/testify/require"

//type MockService struct {
//	mock.Mock
//}
//
//func (m *MockService) UpdateMetrics(metricType, metricName, metricValue string) error {
//	args := m.Called(metricType, metricName, metricValue)
//	return args.Error(0)
//}
//
//func (m *MockService) GetMetricValue(metricType, metricName string) (string, error) {
//	args := m.Called(metricType, metricName)
//	return args.String(0), args.Error(1)
//}
//
//func (m *MockService) GetAll() map[string]any {
//	args := m.Called()
//	return args.Get(0).(map[string]any)
//}
//
//func TestMetrics_getMetricValue_found(t *testing.T) {
//	successResponse := "123"
//	successContentType := "text/plain"
//
//	handler := http.NewServeMux()
//
//	mockService := new(MockService)
//
//	mockService.On("GetMetricValue", mock.Anything, mock.Anything).Return(successResponse, nil)
//
//	newMetricsRoutes(handler, mockService)
//
//	srv := httptest.NewServer(handler)
//	defer srv.Close()
//
//	client := srv.Client()
//
//	t.Run("return metric value", func(t *testing.T) {
//		res, err := client.Get(srv.URL + "/value/abcd/abcd")
//		require.NoError(t, err)
//
//		resBody, err := io.ReadAll(res.Body)
//		require.NoError(t, err)
//
//		defer res.Body.Close()
//
//		assert.Equal(t, http.StatusOK, res.StatusCode)
//		assert.Equal(t, successContentType, res.Header.Get("content-type"))
//		assert.Equal(t, []byte(successResponse), resBody)
//	})
//}
//
//func TestMetrics_getMetricValue_not_found(t *testing.T) {
//	handler := http.NewServeMux()
//
//	mockService := new(MockService)
//
//	mockService.On("GetMetricValue", mock.Anything, mock.Anything).Return("", errors.New("Metric not found"))
//
//	newMetricsRoutes(handler, mockService)
//
//	srv := httptest.NewServer(handler)
//	defer srv.Close()
//
//	client := srv.Client()
//
//	t.Run("metric not found", func(t *testing.T) {
//		res, err := client.Get(srv.URL + "/value/abcd/abcd")
//		require.NoError(t, err)
//
//		_, err = io.Copy(io.Discard, res.Body)
//		require.NoError(t, err)
//
//		defer res.Body.Close()
//
//		assert.Equal(t, http.StatusNotFound, res.StatusCode)
//	})
//}
//
//func TestMetrics_updateMetrics(t *testing.T) {
//	handler := http.NewServeMux()
//
//	mockService := new(MockService)
//
//	mockService.On("UpdateMetrics", mock.Anything, mock.Anything, mock.Anything).Return(nil)
//
//	newMetricsRoutes(handler, mockService)
//
//	srv := httptest.NewServer(handler)
//	defer srv.Close()
//
//	client := srv.Client()
//
//	type request struct {
//		method      string
//		contentType string
//	}
//
//	tests := []struct {
//		name         string
//		request      request
//		expectedCode int
//	}{
//		{
//			name: "GET method not allowed",
//			request: request{
//				method: http.MethodGet,
//			},
//			expectedCode: http.StatusMethodNotAllowed,
//		},
//		{
//			name: "PATCH method not allowed",
//			request: request{
//				method: http.MethodPatch,
//			},
//			expectedCode: http.StatusMethodNotAllowed,
//		},
//		{
//			name: "PUT method not allowed",
//			request: request{
//				method: http.MethodPut,
//			},
//			expectedCode: http.StatusMethodNotAllowed,
//		},
//		{
//			name: "DELETE method not allowed",
//			request: request{
//				method: http.MethodDelete,
//			},
//			expectedCode: http.StatusMethodNotAllowed,
//		},
//		{
//			name: "bad content type",
//			request: request{
//				method:      http.MethodPost,
//				contentType: "application/json",
//			},
//			expectedCode: http.StatusBadRequest,
//		},
//		{
//			name: "successful case",
//			request: request{
//				method:      http.MethodPost,
//				contentType: "text/plain",
//			},
//			expectedCode: http.StatusOK,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			req, err := http.NewRequest(tt.request.method, srv.URL+"/update/gauge/name/123", nil)
//			require.NoError(t, err)
//
//			req.Header.Set("content-type", tt.request.contentType)
//
//			res, err := client.Do(req)
//			require.NoError(t, err)
//
//			_, err = io.Copy(io.Discard, res.Body)
//			require.NoError(t, err)
//
//			defer res.Body.Close()
//
//			assert.Equal(t, tt.expectedCode, res.StatusCode)
//		})
//	}
//}
//
//func TestMetrics_updateMetrics_notFound(t *testing.T) {
//	handler := http.NewServeMux()
//
//	mockService := new(MockService)
//
//	mockService.On("UpdateMetrics", mock.Anything, mock.Anything, mock.Anything).Return(services.ErrMetricsEmptyName)
//
//	newMetricsRoutes(handler, mockService)
//
//	srv := httptest.NewServer(handler)
//	defer srv.Close()
//
//	client := srv.Client()
//
//	t.Run("metric name is empty", func(t *testing.T) {
//		req, err := http.NewRequest(http.MethodPost, srv.URL+"/update/gauge/name/123", nil)
//		require.NoError(t, err)
//
//		req.Header.Set("content-type", "text/plain")
//
//		res, err := client.Do(req)
//		require.NoError(t, err)
//
//		_, err = io.Copy(io.Discard, res.Body)
//		require.NoError(t, err)
//
//		defer res.Body.Close()
//
//		assert.Equal(t, http.StatusNotFound, res.StatusCode)
//	})
//}
//
//func TestMetrics_updateMetrics_badRequest(t *testing.T) {
//	handler := http.NewServeMux()
//
//	mockService := new(MockService)
//
//	mockService.On("UpdateMetrics", mock.Anything, mock.Anything, mock.Anything).Return(services.ErrMetricsUnknownType)
//
//	newMetricsRoutes(handler, mockService)
//
//	srv := httptest.NewServer(handler)
//	defer srv.Close()
//
//	client := srv.Client()
//
//	t.Run("metric type is unknown", func(t *testing.T) {
//		req, err := http.NewRequest(http.MethodPost, srv.URL+"/update/gauge/name/123", nil)
//		require.NoError(t, err)
//
//		req.Header.Set("content-type", "text/plain")
//
//		res, err := client.Do(req)
//		require.NoError(t, err)
//
//		_, err = io.Copy(io.Discard, res.Body)
//		require.NoError(t, err)
//
//		defer res.Body.Close()
//
//		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
//	})
//}
