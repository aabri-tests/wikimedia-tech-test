package handlers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/wikimedia/internal/infrastructure/adapters/api/http/handlers"
	"github.com/wikimedia/internal/mocks"
)

func TestHealthCheckHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	healthCheck := mocks.NewMockWikiMediaUseCase(ctrl)
	httpHandler := handlers.New(healthCheck)
	r := gin.Default()
	r.GET("/health", httpHandler.HealthCheck)

	// Create a mock HTTP request
	req, _ := http.NewRequest("GET", "/health", nil)

	// Create a mock HTTP response recorder
	w := httptest.NewRecorder()

	// Send the mock request to the router
	r.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, w.Code, http.StatusOK)

	// Check the response body
	expectedBody := `{"status":"ok"}`
	assert.Equal(t, w.Body.String(), expectedBody)
}
func TestSearchHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	keyword := "foo"
	language := "en"
	search := mocks.NewMockWikiMediaUseCase(ctrl)
	search.EXPECT().Search(keyword, language).AnyTimes().Return("Bar", nil)
	httpHandler := handlers.New(search)
	r := gin.Default()
	r.GET("/search", httpHandler.Search)

	// Create a mock HTTP request
	req, _ := http.NewRequest("GET", fmt.Sprintf("/search?keyword=%s", keyword), nil)

	// Create a mock HTTP response recorder
	w := httptest.NewRecorder()

	// Send the mock request to the router
	r.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, w.Code, http.StatusOK)

	// Check the response body
	expectedBody := `{"name":"Foo","short_description":"Bar","language":"en","status":{"code":200,"message":"Ok"}}`
	assert.Equal(t, w.Body.String(), expectedBody)
}
func TestSearchHandlerWithValidationError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	keyword := ""
	language := "en"
	search := mocks.NewMockWikiMediaUseCase(ctrl)
	search.EXPECT().Search(keyword, language).AnyTimes().Return("Bar", nil)
	httpHandler := handlers.New(search)
	r := gin.Default()
	r.GET("/search", httpHandler.Search)

	// Create a mock HTTP request
	req, _ := http.NewRequest("GET", fmt.Sprintf("/search?keyword=%s", keyword), nil)

	// Create a mock HTTP response recorder
	w := httptest.NewRecorder()

	// Send the mock request to the router
	r.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, w.Code, http.StatusBadRequest)
}
func TestMetricsHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	metrics := mocks.NewMockWikiMediaUseCase(ctrl)
	httpHandler := handlers.New(metrics)
	r := gin.Default()
	r.GET("/metrics", httpHandler.Metrics())

	// Create a mock HTTP request
	req, _ := http.NewRequest("GET", "/metrics", nil)

	// Create a mock HTTP response recorder
	w := httptest.NewRecorder()

	// Send the mock request to the router
	r.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, w.Code, http.StatusOK)

}
