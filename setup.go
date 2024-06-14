package main

import (
	"context"
	"net/http"
	"resto_go/server"
	"resto_go/service"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

func serviceSetup(logger *zap.Logger) *service.Service {
	connStr := "postgresql://romi:romi@172.18.0.1:5432/postgres?sslmode=disable"
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		logger.Fatal("Error parsing config", zap.Error(err))
	}
	ctx := context.Background()
	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		logger.Fatal("Error connecting to database", zap.Error(err))
	}

	svc := service.NewService(logger, pool)
	return &svc
}

func serverSetup(log *zap.Logger, svc service.Service) *server.Server {
	s := server.NewServer(log, svc, svc.Pool())
	http.HandleFunc("/getIDs", s.GetIDsHandler)

	// start server
	port := ":8080"
	log.Sugar().Infof("Listening port: %s", port)
	log.Sugar().Fatal(http.ListenAndServe(port, nil))

	return s

}
