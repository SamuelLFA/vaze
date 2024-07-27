package http

import (
	"net/http"
)

type HealthHandler interface {
	HealthCheck(w http.ResponseWriter, r *http.Request)
}

type UserHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
}

func RegisterAPIs(healthHandler HealthHandler, userHandler UserHandler) {
	http.HandleFunc("/health", healthHandler.HealthCheck)
	http.HandleFunc("/v1/users", userHandler.Create)
}
