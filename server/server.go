package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"resto_go/service"
	"resto_go/types"
	u "resto_go/utils"
	"strconv"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type Server struct {
	logger *zap.Logger
	Svc    service.Service
	ctx    context.Context
	pool   *pgxpool.Pool
}

var (
	errInvalidParams      = errors.New("latitude & longitude must be valid numbers")
	errMissingFields      = errors.New("latitude & longitude are required, non-zero and valid")
	errServiceUnavailable = errors.New("service unavailable")
)

func NewServer(logger *zap.Logger, svc service.Service, pool *pgxpool.Pool) *Server {
	return &Server{
		logger: logger,
		Svc:    svc,
		pool:   pool,
		ctx:    context.Background(),
	}
}

func (s *Server) GetIDsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(types.ErrorResponse{Error: "method_not_allowed", Cause: "only GET is allowed"})
		return
	}

	idsChan := make(chan types.Output)
	errChan := make(chan error)

	latStr := r.URL.Query().Get("latitude")
	lonStr := r.URL.Query().Get("longitude")

	in, err := ValidateInputData(latStr, lonStr)
	if err != nil {
		status := http.StatusUnprocessableEntity
		errorCode := "unprocessable_entity"
		if errors.Is(err, errInvalidParams) {
			status = http.StatusBadRequest
			errorCode = "bad_request"
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(types.ErrorResponse{Error: errorCode, Cause: err.Error()})
		return
	}

	go func() {
		info, err := u.GetMerchants(s.ctx, s.pool)
		if err != nil {
			if err.Error() == "no merchants found" {
				if procErr := s.ProcessFile("template/csv_info.csv", s.pool); procErr != nil {
					errChan <- fmt.Errorf("%w: %v", errServiceUnavailable, procErr)
					return
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
		} else if errors.Is(err, errServiceUnavailable) {
			status = http.StatusServiceUnavailable
			cause = "service_unavailable"
		} else {
			status = http.StatusInternalServerError
			cause = "internal_server_error"
		}
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(types.ErrorResponse{Cause: err.Error(), Error: cause})
	}
}

func ValidateInputData(latStr, lonStr string) (types.InputData, error) {
	if latStr == "" || lonStr == "" {
		return types.InputData{}, errMissingFields
	}

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		return types.InputData{}, errInvalidParams
	}

	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		return types.InputData{}, errInvalidParams
	}

	if lat == 0 || lon == 0 {
		return types.InputData{}, errMissingFields
	}

	return types.InputData{Latitude: lat, Longitude: lon}, nil
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
