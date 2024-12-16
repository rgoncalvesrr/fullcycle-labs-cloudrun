package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/rgoncalvesrr/fullcycle-labs-cloudrun/adapter"
	"github.com/rgoncalvesrr/fullcycle-labs-cloudrun/application"
	"github.com/rgoncalvesrr/fullcycle-labs-cloudrun/configs"
	"net/http"
)

type ResultError struct {
	Message string `json:"message"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{cep}", handleRequest)

	_ = http.ListenAndServe(":8080", mux)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	cep := r.PathValue("cep")
	cfg := configs.NewConfig()
	coordinateService := adapter.NewCoordinateServiceAdapter(cfg)
	weatherService := adapter.NewWeatherServiceAdapter(cfg)
	w.Header().Set("Content-Type", "application/json")
	s := application.NewWeatherHandler(coordinateService, weatherService)
	output, e := s.GetTemperature(context.Background(), cep)

	if e != nil {
		switch {
		case errors.Is(e, application.ErrCepInvalid), errors.Is(e, application.ErrCepMalformed):
			w.WriteHeader(http.StatusUnprocessableEntity)
			_ = json.NewEncoder(w).Encode(&ResultError{Message: "invalid zipcode"})
		case errors.Is(e, application.ErrCepNotFound):
			w.WriteHeader(http.StatusNotFound)
			_ = json.NewEncoder(w).Encode(&ResultError{Message: "can not find zipcode"})
		default:
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(&ResultError{Message: "internal server error" + e.Error()})
		}
		return
	}

	_ = json.NewEncoder(w).Encode(output)
}
