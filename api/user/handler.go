package user

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/SamuelLFA/vaze/api/shared"
	"github.com/SamuelLFA/vaze/api/util"
)

type userHandler struct{}

func NewHandler() *userHandler {
	return &userHandler{}
}

func (*userHandler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		util.BuildJSONResponse(
			w,
			shared.ErrorResponse{ErrorMessage: "Unable to read body"},
			http.StatusBadRequest,
		)
		return
	}
	defer r.Body.Close()

	var userRequest userRequest

	err = json.Unmarshal(body, &userRequest)
	if err != nil {
		util.BuildJSONResponse(
			w,
			shared.ErrorResponse{ErrorMessage: "Unable to parse JSON"},
			http.StatusBadRequest,
		)
		return
	}

	if err = userRequest.validate(); err != nil {
		util.BuildJSONResponse(
			w,
			shared.ErrorResponse{ErrorMessage: err.Error()},
			http.StatusBadRequest,
		)
		return
	}

	util.BuildJSONResponse(w, userRequest, http.StatusCreated)
}

type userRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *userRequest) validate() error {
	if u.Email == "" || !isValidEmail(u.Email) {
		return fmt.Errorf("invalid email: %s", u.Email)
	}

	if u.Password == "" || len(u.Password) < 8 || len(u.Password) > 128 {
		return fmt.Errorf("password must be a length between 8 and 128 characteres")
	}

	return nil
}

func isValidEmail(email string) bool {
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
