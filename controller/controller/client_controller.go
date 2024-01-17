package controller

import (
	"encoding/json"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/controller/serializer"
	"log/slog"
	"net/http"
	"strconv"

	dto "github.com/ViniAlvesMartins/tech-challenge-fiap/controller/serializer/input"
)

type ClientController struct {
	clientService contract.ClientUseCase
	logger        *slog.Logger
}

func NewClientController(clientService contract.ClientUseCase, logger *slog.Logger) *ClientController {
	return &ClientController{
		clientService: clientService,
		logger:        logger,
	}
}

func (c *ClientController) CreateClient(w http.ResponseWriter, r *http.Request) {
	var clientDto dto.ClientDto
	var response Response

	err := json.NewDecoder(r.Body).Decode(&clientDto)

	if err != nil {
		c.logger.Error("Unable to decode the request body.  %v", err)
	}

	errValidate := serializer.Validate(clientDto)

	if len(errValidate.Errors) > 0 {
		c.logger.Error("validate error", slog.Any("error", errValidate))
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errValidate)
		return
	}

	validClient, errValidClient := c.clientService.GetAlreadyExists(clientDto.Cpf, clientDto.Email)

	if errValidClient != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if validClient != nil {
		response = Response{
			MessageError: "Client already exists",
			Data:         nil,
		}
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(response)
		return
	}

	clientDomain := clientDto.ConvertEntity()

	client, err := c.clientService.Create(clientDomain)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response = Response{
		MessageError: "",
		Data:         client,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

func (c *ClientController) GetClientByCpf(w http.ResponseWriter, r *http.Request) {
	cpf := r.URL.Query().Get("cpf")
	var response Response

	cpfInt, err := strconv.Atoi(cpf)

	if err != nil {
		c.logger.Error("Error to convert cpf to int.  %v", err)
	}

	client, err := c.clientService.GetClientByCpf(cpfInt)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if client == nil {
		response = Response{
			MessageError: "Not found",
			Data:         nil,
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	response = Response{
		MessageError: "",
		Data:         client,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}

}
