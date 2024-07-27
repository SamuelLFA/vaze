package http

import (
	"net/http"
)

type HealthHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
}

func RegisterAPIs(health HealthHandler) {
	http.HandleFunc("/health", health.Get)
}
