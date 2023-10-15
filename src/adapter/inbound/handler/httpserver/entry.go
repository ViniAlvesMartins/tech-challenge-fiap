package httpserver

import (
	"context"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/adapter/inbound/controller"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
)

func Run(ctx context.Context, logger *slog.Logger, db *gorm.DB) error {
	router := mux.NewRouter()

	userController := controller.NewUserController(logger)

	router.HandleFunc("/users", Chain(func(w http.ResponseWriter, r *http.Request) { userController.GetUser(w, r) }, Method("GET"), Logging()))

	return http.ListenAndServe(":8080", router)
}
