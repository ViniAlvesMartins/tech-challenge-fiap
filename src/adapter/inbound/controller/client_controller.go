package controller

import (
	"encoding/json"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/port"
	"log/slog"
	"net/http"
	"strconv"
)

type ClientController struct {
	clientService port.ClientService
	logger        *slog.Logger
}

func NewClientController(clientService port.ClientService, logger *slog.Logger) *ClientController {
	return &ClientController{
		clientService: clientService,
		logger:        logger,
	}
}

func (c *ClientController) CreateClient(w http.ResponseWriter, r *http.Request) {

	var client domain.Client

	err := json.NewDecoder(r.Body).Decode(&client)

	if err != nil {
		c.logger.Error("Unable to decode the request body.  %v", err)
	}

	clientCreated, err := c.clientService.Create(client.Cpf, client.Name, client.Email)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(clientCreated)
	if err != nil {
		return
	}
}

func (c *ClientController) GetClientByCpf(w http.ResponseWriter, r *http.Request) {
	cpf := r.URL.Query().Get("cpf")

	cpfInt, err := strconv.Atoi(cpf)

	if err != nil {
		c.logger.Error("Error to convert cpf to int.  %v", err)
	}

	client, err := c.clientService.GetClientByCpf(cpfInt)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(client)
	if err != nil {
		return
	}

}
