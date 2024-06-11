package main

import (
	"net/http"

	"go.uber.org/zap"
)

func setup(logger *zap.Logger) error {

	http.HandleFunc("/getIDs", func(w http.ResponseWriter, r *http.Request) {
		getIDsHandler(w, r)
	})

	// start server
	port := ":8080"
	logger.Sugar().Infof("Listening port:", port)
	logger.Sugar().Fatal(http.ListenAndServe(port, nil))

	return nil
}
