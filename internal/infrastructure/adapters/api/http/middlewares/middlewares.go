package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/wikimedia/internal/application/ports"
)

func ProvideMiddleware(_ ports.LogInfoFormat) *[]gin.HandlerFunc {
	return &[]gin.HandlerFunc{
		// Add more middlewares to this list
	}
}
