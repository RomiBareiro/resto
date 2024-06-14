package main

import (
	"net/http"
	"resto_go/server"
	"resto_go/service"

	"go.uber.org/zap"
)

func serviceSetup(log *zap.Logger) *service.Service {
	svc := service.NewService(log)
	//Here we should configure db if we want to do the improvements
	return &svc
}

func serverSetup(log *zap.Logger, svc service.Service) *server.Server {
	s := server.NewServer(log, svc)
	http.HandleFunc("/getIDs", s.GetIDsHandler)

	// start server
	port := ":8080"
	log.Sugar().Infof("Escuchando en el puerto %s", port)
	log.Sugar().Fatal(http.ListenAndServe(port, nil))

	return s

}
