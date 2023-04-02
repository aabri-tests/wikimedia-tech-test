package middlewares_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/wikimedia/internal/infrastructure/adapters/api/http/middlewares"
	"github.com/wikimedia/internal/mocks"
)

func TestProvideMiddleware(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock the logger dependency
	logger := mocks.NewMockLogInfoFormat(ctrl)

	// Call the function to get the middleware handlers
	handlers := middlewares.ProvideMiddleware(logger)

	// Check if the returned value is a non-nil array with at least one middleware handler
	assert.NotNil(t, handlers)
	assert.True(t, len(*handlers) == 0)
}
