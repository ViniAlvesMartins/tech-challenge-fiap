package controller

import (
	"encoding/json"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/entity"
	"log/slog"
	"net/http"
)

type UserController struct {
	logger *slog.Logger
}

func NewUserController(logger *slog.Logger) *UserController {
	return &UserController{
		logger: logger,
	}
}

func (u *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	peter := &entity.User{
		Firstname: "John",
		Lastname:  "Doe",
		Age:       25,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(peter)

	return
}
