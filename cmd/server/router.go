package server

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/wikimedia/docs"
)

func (ds *Server) MapRoutes() {
	ds.router.GET("search", ds.httpHandler.Search)
	ds.router.GET("/health", ds.httpHandler.HealthCheck)
	ds.router.GET("/metrics", ds.httpHandler.Metrics())
	ds.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
