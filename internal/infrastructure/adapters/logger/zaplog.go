package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	Logger *zap.SugaredLogger
}

func (l *Logger) Debug(args ...interface{}) {
	l.Logger.Debug(args)
}

func (l *Logger) Info(args ...interface{}) {
	l.Logger.Info(args)
}

func (l *Logger) Warn(args ...interface{}) {
	l.Logger.Warn(args)
}

func (l *Logger) Error(args ...interface{}) {
	l.Logger.Error(args)
}

func (l *Logger) Panic(args ...interface{}) {
	l.Logger.Panic(args)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.Logger.Fatal(args)
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.Logger.Debugf(template, args)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.Logger.Infof(template, args)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	l.Logger.Warnf(template, args)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.Logger.Errorf(template, args)
}

func (l *Logger) Panicf(template string, args ...interface{}) {
	l.Logger.Panicf(template, args)
}

func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.Logger.Fatalf(template, args)
}
