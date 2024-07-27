package health

import (
	"net/http"

	"github.com/SamuelLFA/vaze/api/util"
)

type healthHandler struct{}

type getResponse struct {
	Ok bool `json:"ok"`
}

func New() *healthHandler {
	return &healthHandler{}
}

func (*healthHandler) Get(w http.ResponseWriter, r *http.Request) {
	response := getResponse{Ok: true}
	util.BuildJSONResponse(w, response, http.StatusOK)
}
