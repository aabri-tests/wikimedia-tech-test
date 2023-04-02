package zaplogger

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/wikimedia/internal/application/ports"
	logger2 "github.com/wikimedia/internal/infrastructure/adapters/logger"
	"github.com/wikimedia/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// EnvNotSupported is an error that is returned when the logger environment specified in the configuration object is not supported.
var EnvNotSupported = errors.New("logger environment not supported.")

// ZapLoggerFactory is a struct that implements the ports.LoggerFactory interface and provides a factory for creating a logger that uses the zap logging library.
type ZapLoggerFactory struct {
	Cfg *config.Config
}

// New is a factory function that returns an instance of the ZapLoggerFactory struct.
func New(c *config.Config) ports.LoggerFactory {
	return &ZapLoggerFactory{
		Cfg: c,
	}
}

// GetLogger is a method of the ZapLoggerFactory struct that returns an instance of the logger.
func (z *ZapLoggerFactory) GetLogger() (ports.LogInfoFormat, error) {
	// Create a new instance of the zap logger.
	zl, er := NewZapLogger(z.Cfg)
	if er != nil {
		// If there was an error creating the logger, return an error.
		return nil, errors.Wrap(er, "zaplogger.GetLogger")
	}
	// Return the logger wrapped in a Logger object.
	return &logger2.Logger{Logger: zl}, nil
}

// NewZapLogger is a function that creates a new instance of the zap logger, based on the configuration passed in as an argument.
func NewZapLogger(config *config.Config) (*zap.SugaredLogger, error) {
	loggerConfig := buildLoggerConfig(config)

	// Build the logger and return it.
	log, err := loggerConfig.Build()
	if err != nil {
		return nil, errors.Wrap(err, "zaplogger.NewZapLogger")
	}
	defer log.Sync()

	return log.Sugar(), nil
}

// buildLoggerConfig builds the logger configuration based on the given config object
func buildLoggerConfig(config *config.Config) zap.Config {
	loggerConfig := zap.Config{}

	switch config.Logger.Environment {
	case "dev", "development":
		loggerConfig = zap.NewDevelopmentConfig()
	case "prod", "production":
		loggerConfig = zap.NewProductionConfig()
		loggerConfig.DisableStacktrace = true
		loggerConfig.Encoding = "json"
		loggerConfig.OutputPaths = []string{config.Logger.Filename}
	default:
		panic(EnvNotSupported)
	}

	loggerConfig.Level = zap.NewAtomicLevelAt(getLevel(config.Logger.Loglevel))

	return loggerConfig
}

// getLevel is a function that converts a string log level to a zapcore.Level type.
func getLevel(level string) zapcore.Level {
	var zapLevel zapcore.Level
	switch strings.ToLower(level) {
	case "debug":
		zapLevel = zap.DebugLevel

	case "info":
		zapLevel = zap.InfoLevel

	case "warn", "warning":
		zapLevel = zap.WarnLevel

	case "error":
		zapLevel = zap.ErrorLevel
	}
	return zapLevel
}
