package main

import (
	"resto_go/service"

	"go.uber.org/zap"
)

func setup(log *zap.Logger) service.Service {
	svc := service.NewService(log)
	//Here we should configure db if we want to do the improvements
	return svc
}
