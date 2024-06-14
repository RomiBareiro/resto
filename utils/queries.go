package utils

import (
	"context"
	"fmt"
	t "resto_go/types"
	"strings"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func UpsertMerchants(ctx context.Context, db *pgxpool.Pool, merchants []t.MerchantInfo) error {
	fmt.Printf("UpsertMerchants")
	conn, err := db.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("could not acquire connection: %v", err)
	}
	defer conn.Release()

	for _, merchant := range merchants {
		query := fmt.Sprintf(`
			INSERT INTO restaurante.merchants (id, latitude, longitude, availability_radius, open_hour, close_hour, rating)
			VALUES ('%s', %v, %v, %v, '%v', '%v', %d)
			ON CONFLICT (id) DO UPDATE
			SET latitude = EXCLUDED.latitude,
				longitude = EXCLUDED.longitude,
				availability_radius = EXCLUDED.availability_radius,
				open_hour = EXCLUDED.open_hour,
				close_hour = EXCLUDED.close_hour,
				rating = EXCLUDED.rating
		`, merchant.ID, merchant.Latitude, merchant.Longitude, merchant.AvailabilityRadius, merchant.OpenHour, merchant.CloseHour, merchant.Rating)
		_, err = conn.Exec(ctx, query)
		if err != nil {
			return fmt.Errorf("could not execute upsert: %v", err)
		}
	}

	return nil
}

func GetMerchants(ctx context.Context, db *pgxpool.Pool) ([]t.MerchantInfo, error) {
	var merchants []t.MerchantInfo

	rows, err := db.Query(ctx, "SELECT * FROM restaurante.merchants")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var data t.MerchantInfo
		err := rows.Scan(&data.ID, &data.Latitude, &data.Longitude, &data.AvailabilityRadius, &data.OpenTime, &data.CloseTime, &data.Rating)
		if err != nil {
			return nil, err
		}
		merchants = append(merchants, data)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return merchants, nil
}

func parseTime(s string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05+00", strings.TrimSpace(s))
}
