package log

import "go.uber.org/zap"

var logger *zap.Logger

func init() {
	logger, _ = zap.NewDevelopment()
}

func Log() *zap.Logger {
	return logger
}
