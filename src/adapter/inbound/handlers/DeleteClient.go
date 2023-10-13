package handlers

import (
	"encoding/json"
	"fiappos/ViniAlvesMartins/tech-challenge-fiap/src/adapter/outbound/repository/entity"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h handler) DeleteClient(w http.ResponseWriter, r *http.Request) {
	// Read the dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Find the book by Id

	var clientEntity entity.ClientEntity

	if result := h.DB.First(&clientEntity, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	h.DB.Delete(&clientEntity)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}
