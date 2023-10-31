package controller

import (
	"encoding/json"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/adapter/inbound/dto"
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
	var clientDto dto.ClientDto

	err := json.NewDecoder(r.Body).Decode(&clientDto)

	if err != nil {
		c.logger.Error("Unable to decode the request body.  %v", err)
	}

	errValidate := dto.Validate(clientDto)

	if len(errValidate.Errors) > 0 {
		c.logger.Error("validate error", slog.Any("error", errValidate))
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errValidate)
		return
	}

	clientDomain := clientDto.ConvertEntity()

	client, err := c.clientService.Create(clientDomain)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(client)
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
