package user

import (
	"net/http"

	"github.com/SamuelLFA/vaze/api/util"
)

type userHandler struct{}

func NewHandler() *userHandler {
	return &userHandler{}
}

func (*userHandler) Create(w http.ResponseWriter, r *http.Request) {
	response := struct {
		Test string
	}{
		Test: "OK",
	}
	util.BuildJSONResponse(w, response, http.StatusCreated)
}
