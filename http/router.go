package http

import (
	"net/http"

	"github.com/SamuelLFA/vaze/api/shared"
	"github.com/SamuelLFA/vaze/api/util"
)

type HealthHandler interface {
	HealthCheck(w http.ResponseWriter, r *http.Request)
}

type UserHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
}

func validateMethod(
	method string,
	handler func(w http.ResponseWriter, r *http.Request),
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			response := shared.ErrorResponse{
				ErrorMessage: "Method not allowed",
			}
			util.BuildJSONResponse(w, response, http.StatusMethodNotAllowed)
			return
		}

		handler(w, r)
	}
}

func RegisterAPIs(healthHandler HealthHandler, userHandler UserHandler) {
	http.HandleFunc("/health", validateMethod("GET", healthHandler.HealthCheck))
	http.HandleFunc("/v1/users", validateMethod("POST", userHandler.Create))
}
