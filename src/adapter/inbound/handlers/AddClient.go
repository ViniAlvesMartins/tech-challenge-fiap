package handlers

import (
	"encoding/json"
	"fiappos/ViniAlvesMartins/tech-challenge-fiap/src/adapter/outbound/repository/entity/converter"
	core "fiappos/ViniAlvesMartins/tech-challenge-fiap/src/core/domain"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (h handler) AddClient(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var client core.ClientDomain
	json.Unmarshal(body, &client)

	clientEntity := converter.ConvertDomainToEntity(&client)

	if result := h.DB.Create(&clientEntity); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}
