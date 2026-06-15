package user_transport_http

import (
	"encoding/json"
	"fmt"
	"net/http"

	core_logger "github.com/Vidgimka/eventi/internal/core/logger"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

type CreateUserResponse struct {
	ID          string  `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func (h *UserHTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	log.Debug("invoce  CreateUser handler")

	var responce CreateUserResponse
	if err := json.NewDecoder(r.Body).Decode(&responce); err != nil {
		fmt.Println("Error")
	}
}
