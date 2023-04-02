package ports

import (
	"github.com/gin-gonic/gin"
)

type WikiMediaHTTPHandler interface {
	// Search This method is responsible for handling search requests.
	Search(c *gin.Context)
	// HealthCheck This method is responsible for handling health check requests.
	HealthCheck(c *gin.Context)
	// Metrics This method is responsible for handling prometheus metrics requests.
	Metrics() gin.HandlerFunc
}
