package logger_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wikimedia/internal/infrastructure/adapters/logger"
	"go.uber.org/zap/zaptest"
)

func TestLogger(t *testing.T) {
	ts := newTestLogSpy(t)
	defer ts.AssertPassed()

	log := zaptest.NewLogger(ts)
	lg := logger.Logger{
		Logger: log.Sugar(),
	}

	lg.Info("received work order")
	lg.Debug("starting work")
	lg.Warn("work may fail")

	assert.Panics(t, func() {
		lg.Panic("failed to do work")
	}, "log.Panic should panic")

	ts.AssertMessages(
		"INFO	received work order",
		"DEBUG	starting work",
		"WARN	work may fail",
		"PANIC	failed to do work",
	)
}

func TestLoggerWithFormat(t *testing.T) {
	ts := newTestLogSpy(t)
	defer ts.AssertPassed()

	log := zaptest.NewLogger(ts)
	lg := logger.Logger{
		Logger: log.Sugar(),
	}

	lg.Infof("received work order %s", "foo")
	lg.Debugf("starting work %s", "foo")
	lg.Warnf("work may fail %s", "foo")

	assert.Panics(t, func() {
		lg.Panicf("failed to do work %s", "foo")
	}, "log.Panic should panic")

	ts.AssertMessages(
		"INFO	received work order foo",
		"DEBUG	starting work foo",
		"WARN	work may fail foo",
		"PANIC	failed to do work foo",
	)
}
