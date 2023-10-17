package controller

import (
	"encoding/json"
	"fiappos/ViniAlvesMartins/tech-challenge-fiap/src/core/domain"
	"fiappos/ViniAlvesMartins/tech-challenge-fiap/src/core/port"
	"log"
	"net/http"
)

type ClientController struct {
	clientService port.ClientService
}

func NewClientController(clientService port.ClientService) *ClientController {
	return &ClientController{
		clientService: clientService,
	}
}

func (c *ClientController) CreateClient(w http.ResponseWriter, r *http.Request) {

	var client domain.Client

	err := json.NewDecoder(r.Body).Decode(&client)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	clientCreated, err := c.clientService.Create(client.Cpf, client.Name, client.Email)

	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(clientCreated)
	if err != nil {
		return
	}
}
