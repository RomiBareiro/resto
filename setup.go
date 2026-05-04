package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"resto_go/server"
	"resto_go/service"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

const appDBName = "romi"

func serviceSetup(logger *zap.Logger) *service.Service {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "postgresql://romi:romi@db:5432/romi?sslmode=disable"
	}

	ctx := context.Background()
	if err := ensureDatabaseExists(ctx, connStr, appDBName); err != nil {
		logger.Fatal("Error ensuring database exists", zap.Error(err))
	}

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		logger.Fatal("Error parsing config", zap.Error(err))
	}

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		logger.Fatal("Error connecting to database", zap.Error(err))
	}

	svc := service.NewService(logger, pool)
	return &svc
}

func ensureDatabaseExists(ctx context.Context, connStr, dbName string) error {
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return err
	}
	config.ConnConfig.Database = "postgres"

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return err
	}
	defer pool.Close()

	_, err = pool.Exec(ctx, fmt.Sprintf("CREATE DATABASE %s", pgx.Identifier{dbName}.Sanitize()))
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "42P04" {
			return nil
		}
		return err
	}
	return nil
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
