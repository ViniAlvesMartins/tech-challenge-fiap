package controller

import (
	"encoding/json"
	"fiappos/ViniAlvesMartins/tech-challenge-fiap/src/core/port"
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

func (c *ClientController) CreateProduct(w http.ResponseWriter, r *http.Request) {

	client, err := c.clientService.Create("123", 11122233344, "teste", "teste@com")

	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(client)
}
