package main

import (
	"context"
	"net/http"
	"os"
	u "resto_go/utils"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	ctx := context.Background()
	svc := serviceSetup(logger)
	_ = serverSetup(logger, *svc)
	ticker := time.NewTicker(6 * time.Hour)
	pool := (*svc).Pool()
	defer ticker.Stop()

	for {
		<-ticker.C // Check file every 6 hours
		err := downloadFile("https://example.com/file.csv", "template/csv_info.csv")
		if err != nil {
			logger.Sugar().Errorf("Error downloading file: %v\n", err)
			continue // Skip processing if file download failed
		}

		logger.Info("File downloaded successfully")

		if err := ProcessFile("template/csv_info.csv", logger, pool, ctx); err != nil {
			logger.Sugar().Errorf("Error processing file: %v\n", err)
		}
	}
	defer pool.Close()
}

func downloadFile(url string, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	return err
}

// ProcessFile loads file data into our db
func ProcessFile(filepath string, log *zap.Logger, DB *pgxpool.Pool, ctx context.Context) error {
	data, err := u.ReadFile(filepath)
	if err != nil {
		log.Sugar().Errorf("could not read file: %s", filepath)
		return err
	}
	log.Sugar().Infof("data: %v", data)
	if err := u.UpsertMerchants(ctx, DB, data); err != nil {
		return err
	}
	log.Sugar().Infof("UpsertMerchants ok")
	return nil
}
