package handlers

import (
	"encoding/json"
	"fiappos/ViniAlvesMartins/tech-challenge-fiap/src/adapter/outbound/repository/entity"
	"fmt"
	"net/http"
)

func (h handler) GetAllClients(w http.ResponseWriter, r *http.Request) {
	var clients []entity.ClientEntity

	if result := h.DB.Find(&clients); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(clients)
}
