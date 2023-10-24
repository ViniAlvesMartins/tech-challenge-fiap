package controller

import (
	"encoding/json"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/adapter/inbound/dto"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/port"
	"log/slog"
	"net/http"
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

	errValidate := dto.ValidateClient(clientDto)

	if len(errValidate.Errors) > 0 {
		c.logger.Error("validate error.  %v", errValidate)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errValidate)
		return
	}

	clientDomain := dto.ConvertClientDtoToDomain(clientDto)

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
