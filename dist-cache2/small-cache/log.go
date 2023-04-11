package small_cache

import "go.uber.org/zap"

var logger *zap.Logger

func init() {
	logger, _ = zap.NewDevelopment()
}

func LogInstance() *zap.Logger {
	return logger
}
