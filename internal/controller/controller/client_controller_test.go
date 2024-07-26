package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/application/contract/mock"
	dto "github.com/ViniAlvesMartins/tech-challenge-fiap/internal/controller/serializer/input"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/controller/serializer/output"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/entity"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

func TestClientController_CreateClient(t *testing.T) {
	t.Run("create client successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		client := entity.Client{
			ID:    1,
			Cpf:   12345678900,
			Name:  "Test Client",
			Email: "testclient@example.com",
		}

		body := dto.ClientDto{Name: client.Name, Email: client.Email, Cpf: client.Cpf}
		jsonBody, _ := json.Marshal(body)
		bodyReader := bytes.NewReader(jsonBody)

		r, _ := http.NewRequest("POST", "/clients", bodyReader)
		w := httptest.NewRecorder()

		loggerMock := slog.New(slog.NewTextHandler(os.Stderr, nil))

		clientUseCaseMock := mock.NewMockClientUseCase(ctrl)
		checkClient := clientUseCaseMock.EXPECT().GetByCpfOrEmail(client.Cpf, client.Email).Return(nil, nil).Times(1)
		clientUseCaseMock.EXPECT().Create(body.ConvertEntity()).Return(&client, nil).Times(1).After(checkClient)

		c := NewClientController(clientUseCaseMock, loggerMock)
		c.CreateClient(w, r)

		jsonResponse, _ := json.Marshal(Response{
			Error: "",
			Data:  output.ClientFromEntity(client),
		})

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Equal(t, string(jsonResponse), string(w.Body.Bytes()))
	})

	t.Run("error decoding body", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		wrongDTO := struct {
			Data string
		}{
			Data: "wrong data",
		}

		loggerMock := slog.New(slog.NewTextHandler(os.Stderr, nil))
		clientUseCaseMock := mock.NewMockClientUseCase(ctrl)

		jsonBody, _ := json.Marshal(wrongDTO)
		bodyReader := bytes.NewReader(jsonBody)

		r, _ := http.NewRequest("POST", "/clients", bodyReader)
		w := httptest.NewRecorder()

		c := NewClientController(clientUseCaseMock, loggerMock)
		c.CreateClient(w, r)

		jsonResponse, _ := json.Marshal(Response{
			Error: "Invalid body, make sure all required fields are sent",
			Data:  nil,
		})

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, string(jsonResponse), string(w.Body.Bytes()))

	})

	t.Run("check if client already exists error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		expectedErr := errors.New("error connecting to database")

		client := entity.Client{
			ID:    1,
			Cpf:   12345678900,
			Name:  "Test Client",
			Email: "testclient@example.com",
		}

		body := dto.ClientDto{Name: client.Name, Email: client.Email, Cpf: client.Cpf}
		jsonBody, _ := json.Marshal(body)
		bodyReader := bytes.NewReader(jsonBody)

		r, _ := http.NewRequest("POST", "/clients", bodyReader)
		w := httptest.NewRecorder()

		loggerMock := slog.New(slog.NewTextHandler(os.Stderr, nil))

		clientUseCaseMock := mock.NewMockClientUseCase(ctrl)
		clientUseCaseMock.EXPECT().GetByCpfOrEmail(client.Cpf, client.Email).Return(nil, expectedErr).Times(1)

		c := NewClientController(clientUseCaseMock, loggerMock)
		c.CreateClient(w, r)

		jsonResponse, _ := json.Marshal(Response{
			Error: "Error validating client",
			Data:  nil,
		})

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(jsonResponse), string(w.Body.Bytes()))

	})

	t.Run("error creating client", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		expectedErr := errors.New("error connecting to database")

		client := entity.Client{
			ID:    1,
			Cpf:   12345678900,
			Name:  "Test Client",
			Email: "testclient@example.com",
		}

		body := dto.ClientDto{Name: client.Name, Email: client.Email, Cpf: client.Cpf}
		jsonBody, _ := json.Marshal(body)
		bodyReader := bytes.NewReader(jsonBody)

		r, _ := http.NewRequest("POST", "/clients", bodyReader)
		w := httptest.NewRecorder()

		loggerMock := slog.New(slog.NewTextHandler(os.Stderr, nil))

		clientUseCaseMock := mock.NewMockClientUseCase(ctrl)
		checkClient := clientUseCaseMock.EXPECT().GetByCpfOrEmail(client.Cpf, client.Email).Return(nil, nil).Times(1)
		clientUseCaseMock.EXPECT().Create(body.ConvertEntity()).Return(nil, expectedErr).Times(1).After(checkClient)

		c := NewClientController(clientUseCaseMock, loggerMock)
		c.CreateClient(w, r)

		jsonResponse, _ := json.Marshal(Response{
			Error: "Error creating client",
			Data:  nil,
		})

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(jsonResponse), string(w.Body.Bytes()))

	})

	t.Run("invalid client error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		client := entity.Client{
			ID:    1,
			Cpf:   12345678900,
			Name:  "Test Client",
			Email: "testclient@example.com",
		}

		body := dto.ClientDto{Name: client.Name, Email: client.Email, Cpf: client.Cpf}
		jsonBody, _ := json.Marshal(body)
		bodyReader := bytes.NewReader(jsonBody)

		r, _ := http.NewRequest("POST", "/clients", bodyReader)
		w := httptest.NewRecorder()

		loggerMock := slog.New(slog.NewTextHandler(os.Stderr, nil))

		clientUseCaseMock := mock.NewMockClientUseCase(ctrl)
		clientUseCaseMock.EXPECT().GetByCpfOrEmail(client.Cpf, client.Email).Return(&client, nil).Times(1)

		c := NewClientController(clientUseCaseMock, loggerMock)
		c.CreateClient(w, r)

		jsonResponse, _ := json.Marshal(Response{
			Error: "Client already exists",
			Data:  nil,
		})

		assert.Equal(t, http.StatusConflict, w.Code)
		assert.Equal(t, string(jsonResponse), string(w.Body.Bytes()))
	})
}

func TestClientController_GetClientByCpf(t *testing.T) {
	t.Run("get client by cpf successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		client := entity.Client{
			ID:    1,
			Cpf:   12345678900,
			Name:  "Test Client",
			Email: "testclient@example.com",
		}

		body := dto.ClientDto{Name: client.Name, Email: client.Email, Cpf: client.Cpf}
		jsonBody, _ := json.Marshal(body)
		bodyReader := bytes.NewReader(jsonBody)
		cpf := strconv.Itoa(client.Cpf)

		r, _ := http.NewRequest("GET", fmt.Sprintf("/clients?cpf=%s", cpf), bodyReader)
		w := httptest.NewRecorder()

		loggerMock := slog.New(slog.NewTextHandler(os.Stderr, nil))

		clientUseCaseMock := mock.NewMockClientUseCase(ctrl)
		clientUseCaseMock.EXPECT().GetByCpf(client.Cpf).Return(&client, nil).Times(1)

		c := NewClientController(clientUseCaseMock, loggerMock)
		c.GetClientByCpf(w, r)

		jsonResponse, _ := json.Marshal(Response{
			Error: "",
			Data:  output.ClientFromEntity(client),
		})

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(jsonResponse), string(w.Body.Bytes()))
	})

	t.Run("error converting cpf", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		client := entity.Client{
			ID:    1,
			Cpf:   12345678900,
			Name:  "Test Client",
			Email: "testclient@example.com",
		}

		body := dto.ClientDto{Name: client.Name, Email: client.Email, Cpf: client.Cpf}
		jsonBody, _ := json.Marshal(body)
		bodyReader := bytes.NewReader(jsonBody)
		cpf := ""

		r, _ := http.NewRequest("GET", fmt.Sprintf("/clients?cpf=%s", cpf), bodyReader)
		w := httptest.NewRecorder()

		loggerMock := slog.New(slog.NewTextHandler(os.Stderr, nil))

		clientUseCaseMock := mock.NewMockClientUseCase(ctrl)

		c := NewClientController(clientUseCaseMock, loggerMock)
		c.GetClientByCpf(w, r)

		jsonResponse, _ := json.Marshal(Response{
			Error: "Make sure document is an int",
			Data:  nil,
		})

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(jsonResponse), string(w.Body.Bytes()))
	})

	t.Run("error getting client by cpf", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		expectedErr := errors.New("error connecting to database")

		client := entity.Client{
			ID:    1,
			Cpf:   12345678900,
			Name:  "Test Client",
			Email: "testclient@example.com",
		}

		body := dto.ClientDto{Name: client.Name, Email: client.Email, Cpf: client.Cpf}
		jsonBody, _ := json.Marshal(body)
		bodyReader := bytes.NewReader(jsonBody)
		cpf := strconv.Itoa(client.Cpf)

		r, _ := http.NewRequest("GET", fmt.Sprintf("/clients?cpf=%s", cpf), bodyReader)
		w := httptest.NewRecorder()

		loggerMock := slog.New(slog.NewTextHandler(os.Stderr, nil))

		clientUseCaseMock := mock.NewMockClientUseCase(ctrl)
		clientUseCaseMock.EXPECT().GetByCpf(client.Cpf).Return(nil, expectedErr).Times(1)

		c := NewClientController(clientUseCaseMock, loggerMock)
		c.GetClientByCpf(w, r)

		jsonResponse, _ := json.Marshal(Response{
			Error: "Error finding client",
			Data:  nil,
		})

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(jsonResponse), string(w.Body.Bytes()))
	})

	t.Run("client not found error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		client := entity.Client{
			ID:    1,
			Cpf:   12345678900,
			Name:  "Test Client",
			Email: "testclient@example.com",
		}

		body := dto.ClientDto{Name: client.Name, Email: client.Email, Cpf: client.Cpf}
		jsonBody, _ := json.Marshal(body)
		bodyReader := bytes.NewReader(jsonBody)
		cpf := strconv.Itoa(client.Cpf)

		r, _ := http.NewRequest("GET", fmt.Sprintf("/clients?cpf=%s", cpf), bodyReader)
		w := httptest.NewRecorder()

		loggerMock := slog.New(slog.NewTextHandler(os.Stderr, nil))

		clientUseCaseMock := mock.NewMockClientUseCase(ctrl)
		clientUseCaseMock.EXPECT().GetByCpf(client.Cpf).Return(nil, nil).Times(1)

		c := NewClientController(clientUseCaseMock, loggerMock)
		c.GetClientByCpf(w, r)

		jsonResponse, _ := json.Marshal(Response{
			Error: "Client not found",
			Data:  nil,
		})

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Equal(t, string(jsonResponse), string(w.Body.Bytes()))
	})
}
