package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"resto_go/service"
	"resto_go/types"
	u "resto_go/utils"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	if err := setup(logger); err != nil {
		logger.Sugar().Errorf("could not get input params: %v", err)
		return
	}

}

func getIDsHandler(w http.ResponseWriter, r *http.Request) {
	idsChan := make(chan types.Output)
	errChan := make(chan error)

	in, err := ValidateInputData(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	go func() {
		info, err := u.ReadFile("template/csv_info.csv")
		if err != nil {
			errChan <- err
			return
		}
		ids, err := service.GetIDS(in, info)
		if err != nil {
			errChan <- err
			return
		}
		idsChan <- ids
	}()

	select {
	case ids := <-idsChan:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ids)
	case err := <-errChan:
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
