package httpserver

import (
	"context"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/adapter/inbound/controller"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/port"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
)

type Entry struct {
	logger        *slog.Logger
	clientService port.ClientService
}

func NewEntry(logger *slog.Logger, clientService port.ClientService) *Entry {
	return &Entry{
		logger:        logger,
		clientService: clientService,
	}
}

func (e *Entry) Run(ctx context.Context) error {
	router := mux.NewRouter()

	userController := controller.NewUserController(e.logger)
	clientController := controller.NewClientController(e.clientService, e.logger)

	router.HandleFunc("/users", Chain(func(w http.ResponseWriter, r *http.Request) { userController.GetUser(w, r) }, Method("GET"), Logging()))
	router.HandleFunc("/client", Chain(func(w http.ResponseWriter, r *http.Request) { clientController.CreateClient(w, r) }, Method("POST"), Logging()))

	return http.ListenAndServe(":8080", router)
}
