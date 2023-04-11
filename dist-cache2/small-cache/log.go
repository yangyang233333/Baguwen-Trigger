package small_cache

import "go.uber.org/zap"

func a() {
	_, _ = zap.NewProduction()
}
