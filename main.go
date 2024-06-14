package main

import (
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	svc := serviceSetup(logger)
	_ = serverSetup(logger, *svc)

}
