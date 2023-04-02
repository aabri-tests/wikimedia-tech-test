package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wikimedia/internal/application/ports"
	"github.com/wikimedia/pkg/config"
)

var (
	DefaultServerHostEnvVar = "SERVER_HOST"
	DefaultServerPortEnvVar = "SERVER_PORT"
	DefaultReadTimeout      = 5 * time.Second
	DefaultWriteTimeout     = 10 * time.Second
	DefaultMaxHeaderBytes   = 1 << 20
)

type Server struct {
	router      *gin.Engine
	cfg         *config.Config
	httpHandler ports.WikiMediaHTTPHandler
}

func NewServer(cfg *config.Config, middleware *[]gin.HandlerFunc, httpHandler ports.WikiMediaHTTPHandler) *Server {
	engine := gin.New()
	server := &Server{
		router:      engine,
		cfg:         cfg,
		httpHandler: httpHandler,
	}
	engine.ForwardedByClientIP = true

	// Add all middlewares to the router
	for _, m := range *middleware {
		server.router.Use(m)
	}
	gin.SetMode(cfg.Logger.Loglevel)

	return server
}

// Start serving the application
func (ds *Server) Start() error {
	host := os.Getenv(DefaultServerHostEnvVar)
	port, _ := strconv.Atoi(os.Getenv(DefaultServerPortEnvVar))
	url := fmt.Sprintf(ds.cfg.Server.Host, host, port)
	s := &http.Server{
		Addr:           url,
		Handler:        ds.router,
		ReadTimeout:    DefaultReadTimeout,
		WriteTimeout:   DefaultWriteTimeout,
		MaxHeaderBytes: DefaultMaxHeaderBytes,
	}
	return s.ListenAndServe()
}
