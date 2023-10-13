package handlers

import (
	"encoding/json"
	"fiappos/ViniAlvesMartins/tech-challenge-fiap/src/adapter/outbound/repository/entity"
	"fiappos/ViniAlvesMartins/tech-challenge-fiap/src/adapter/outbound/repository/entity/converter"
	core "fiappos/ViniAlvesMartins/tech-challenge-fiap/src/core/domain"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h handler) UpdateClient(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var clientUpdate core.ClientDomain
	json.Unmarshal(body, &clientUpdate)

	clientEntityUpdate := converter.ConvertDomainToEntity(&clientUpdate)

	var clientEntity entity.ClientEntity

	if result := h.DB.First(&clientEntity, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	clientEntity.Cpf = clientEntityUpdate.Cpf
	clientEntity.Name = clientEntityUpdate.Name
	clientEntity.Email = clientEntityUpdate.Email

	h.DB.Save(&clientEntity)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}
