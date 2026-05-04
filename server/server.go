package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"resto_go/service"
	"resto_go/types"
	u "resto_go/utils"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type Server struct {
	logger *zap.Logger
	Svc    service.Service
	ctx    context.Context
	pool   *pgxpool.Pool
}

func NewServer(logger *zap.Logger, svc service.Service, pool *pgxpool.Pool) *Server {
	return &Server{
		logger: logger,
		Svc:    svc,
		pool:   pool,
		ctx:    context.Background(),
	}
}

func (s *Server) GetIDsHandler(w http.ResponseWriter, r *http.Request) {
	idsChan := make(chan types.Output)
	errChan := make(chan error)

	in, err := ValidateInputData(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	go func() {
		info, err := u.GetMerchants(s.ctx, s.pool)
		if err != nil {
			if err.Error() == "no merchants found" {
				if err := s.ProcessFile("template/csv_info.csv", s.pool); err != nil {
					errChan <- fmt.Errorf("error processing file: %v", err)
				}
			}
			errChan <- err
			return
		}
		ids, err := s.Svc.GetIDS(in, info)
		if err != nil {
			errChan <- err
			return
		}
		idsChan <- *ids
	}()

	w.Header().Set("Content-Type", "application/json")
	select {
	case ids := <-idsChan:
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ids)
	case err := <-errChan:
		var (
			status int
			cause  string
		)
		if err.Error() == "no available merchants" {
			status = http.StatusNotFound
			cause = "not_found"

		} else {
			status = http.StatusInternalServerError
			cause = "internal_server_error"
		}
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(types.ErrorResponse{Cause: err.Error(), Error: cause})
	}
}

func ValidateInputData(body io.Reader) (types.InputData, error) {
	var in types.InputData
	err := json.NewDecoder(body).Decode(&in)
	if err != nil {
		return types.InputData{}, errors.New("invalid JSON format")
	}

	if in.Latitude == 0 || in.Longitude == 0 {
		return types.InputData{}, errors.New("latitude & longitude are required and must be non-zero")
	}

	return in, nil
}

// ProcessFile loads file data into our db
func (s *Server) ProcessFile(filepath string, DB *pgxpool.Pool) error {
	data, err := u.ReadFile(filepath)
	if err != nil {
		s.logger.Sugar().Errorf("could not read file: %s", filepath)
		return err
	}
	if err := u.UpsertMerchants(s.ctx, DB, data); err != nil {
		return err
	}
	s.logger.Sugar().Infof("UpsertMerchants ok")
	return nil
}
