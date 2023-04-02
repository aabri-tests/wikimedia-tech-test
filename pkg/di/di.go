package di

import (
	"github.com/wikimedia/cmd/server"
	"github.com/wikimedia/internal/application/ports"
	"github.com/wikimedia/internal/application/usecases"
	"github.com/wikimedia/internal/domain/factories/cache/redis"
	zaplogger "github.com/wikimedia/internal/domain/factories/log/zap"
	"github.com/wikimedia/internal/domain/factories/wikimedia"
	"github.com/wikimedia/internal/infrastructure/adapters/api/http/handlers"
	"github.com/wikimedia/internal/infrastructure/adapters/api/http/middlewares"
	wikimedia2 "github.com/wikimedia/internal/infrastructure/adapters/parsers/wikimedia"
	"github.com/wikimedia/pkg/config"
	"go.uber.org/dig"
)

// Create a new dependency injection container using the `dig` package
var container = dig.New()

// BuildContainer Build the container by adding providers for each dependency
func BuildContainer() *dig.Container {
	// Add a provider for the configuration, which is loaded using `config.New()`
	container.Provide(config.New)

	// Add a provider for Parser
	container.Provide(func(cfg *config.Config) (ports.Parser, error) {
		return &wikimedia2.ShortDescriptionParser{Cfg: cfg}, nil
	})
	// Add a provider for WikiMediaFactory
	container.Provide(func(cfg *config.Config) (ports.WikiMediaFactory, error) {
		return &wikimedia.WikiMediaFactory{Cfg: cfg}, nil
	})
	// Add a provider for the logger factory, which is used to create a logger implementation based on the configuration
	container.Provide(func(cfg *config.Config) ports.LoggerFactory {
		return &zaplogger.ZapLoggerFactory{Cfg: cfg}
	})
	// Add a provider for redis cache factory
	container.Provide(func(cfg *config.Config) ports.CacheFactory {
		return &redis.RedisFactory{Cfg: cfg}
	})
	// Add a provider for the log info format, which is used to create a logger instance based on the configuration
	container.Provide(func(Cfg *config.Config) (ports.LogInfoFormat, error) {
		var logger ports.LoggerFactory
		if Cfg.Logger.Use == "zapLogger" {
			logger = zaplogger.New(Cfg)
		}

		log, err := logger.GetLogger()
		if err != nil {
			return nil, err
		}
		return log, nil
	})
	// Add a provider for WikiMediaUseCase
	container.Provide(func(Cfg *config.Config, parser ports.Parser, log ports.LogInfoFormat) (ports.WikiMediaUseCase, error) {
		// Inject Redis Cache to mongodb factory
		redisCache := redis.New(Cfg)
		cache, err := redisCache.GetCache()
		if err != nil {
			return nil, err
		}
		var wiki ports.WikiMediaFactory
		wiki = wikimedia.New(Cfg)

		wikiAdapter, err := wiki.GetWikiMediaAdapter()
		if err != nil {
			return nil, err
		}
		return usecases.New(wikiAdapter, parser, cache, log), nil
	})
	// Add a provider for Http handler
	container.Provide(handlers.New)
	// Add a provider for the middleware, which is used to add middleware to the HTTP server
	container.Provide(middlewares.ProvideMiddleware)
	// Register the constructor function for creating an instance of the HTTP server.
	container.Provide(server.NewServer)

	return container
}
