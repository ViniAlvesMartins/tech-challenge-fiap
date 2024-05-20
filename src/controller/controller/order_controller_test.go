package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/application/contract/mock"
	dto "github.com/ViniAlvesMartins/tech-challenge-fiap/src/controller/serializer/input"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/controller/serializer/output"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/enum"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestOrderController_CreateOrder(t *testing.T) {
	t.Run("create order successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		product := entity.Product{
			ID:          1,
			NameProduct: "Test product",
			Description: "Test product",
			Price:       123.45,
			CategoryId:  1,
			Active:      true,
		}

		client := entity.Client{
			ID:    1,
			Cpf:   12345678900,
			Name:  "Test Client",
			Email: "testclient@example.com",
		}

		body := dto.OrderDto{
			ClientId: &client.ID,
			Products: []struct {
				ID int `json:"id"`
			}{
				{
					ID: 1,
				},
			},
		}

		products := []*entity.Product{&product}

		order := entity.Order{
			ID:          1,
			ClientId:    &client.ID,
			StatusOrder: enum.AWAITING_PAYMENT,
			Amount:      123.45,
			Products:    products,
		}

		jsonBody, _ := json.Marshal(body)
		bodyReader := bytes.NewReader(jsonBody)

		r, _ := http.NewRequest("POST", "/clients", bodyReader)
		w := httptest.NewRecorder()

		loggerMock := slog.New(slog.NewTextHandler(os.Stderr, nil))

		productUseCaseMock := mock.NewMockProductUseCase(ctrl)
		clientUseCaseMock := mock.NewMockClientUseCase(ctrl)
		orderUseCaseMock := mock.NewMockOrderUseCase(ctrl)

		getProduct := productUseCaseMock.EXPECT().GetById(1).Return(&product, nil).Times(1)
		getClient := clientUseCaseMock.EXPECT().GetClientById(body.ClientId).Return(&client, nil).Times(1).After(getProduct)
		orderUseCaseMock.EXPECT().Create(body.ConvertToEntity(), products).Return(&order, nil).Times(1).After(getClient)

		c := NewOrderController(orderUseCaseMock, productUseCaseMock, clientUseCaseMock, loggerMock)
		c.CreateOrder(w, r)

		jsonResponse, _ := json.Marshal(Response{
			Error: "",
			Data:  output.OrderFromEntity(order),
		})

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Equal(t, string(jsonResponse), string(w.Body.Bytes()))
	})

	t.Run("body decoding error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		wrongDTO := "test"

		jsonBody, _ := json.Marshal(wrongDTO)
		bodyReader := bytes.NewReader(jsonBody)

		r, _ := http.NewRequest("POST", "/clients", bodyReader)
		w := httptest.NewRecorder()

		loggerMock := slog.New(slog.NewTextHandler(os.Stderr, nil))

		productUseCaseMock := mock.NewMockProductUseCase(ctrl)
		clientUseCaseMock := mock.NewMockClientUseCase(ctrl)
		orderUseCaseMock := mock.NewMockOrderUseCase(ctrl)

		c := NewOrderController(orderUseCaseMock, productUseCaseMock, clientUseCaseMock, loggerMock)
		c.CreateOrder(w, r)

		jsonResponse, _ := json.Marshal(Response{
			Error: "Unable to decode the request body",
			Data:  nil,
		})

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(jsonResponse), string(w.Body.Bytes()))
	})

	t.Run("empty order error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		body := dto.OrderDto{
			ClientId: nil,
			Products: nil,
		}

		jsonBody, _ := json.Marshal(body)
		bodyReader := bytes.NewReader(jsonBody)

		r, _ := http.NewRequest("POST", "/clients", bodyReader)
		w := httptest.NewRecorder()

		loggerMock := slog.New(slog.NewTextHandler(os.Stderr, nil))

		productUseCaseMock := mock.NewMockProductUseCase(ctrl)
		clientUseCaseMock := mock.NewMockClientUseCase(ctrl)
		orderUseCaseMock := mock.NewMockOrderUseCase(ctrl)

		c := NewOrderController(orderUseCaseMock, productUseCaseMock, clientUseCaseMock, loggerMock)
		c.CreateOrder(w, r)

		jsonResponse, _ := json.Marshal(Response{
			Error: "Product is required",
			Data:  nil,
		})

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, string(jsonResponse), string(w.Body.Bytes()))
	})

	t.Run("error getting product", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		expectedErr := errors.New("error connecting to database")

		body := dto.OrderDto{
			ClientId: nil,
			Products: []struct {
				ID int `json:"id"`
			}{
				{
					ID: 1,
				},
			},
		}

		jsonBody, _ := json.Marshal(body)
		bodyReader := bytes.NewReader(jsonBody)

		r, _ := http.NewRequest("POST", "/clients", bodyReader)
		w := httptest.NewRecorder()

		loggerMock := slog.New(slog.NewTextHandler(os.Stderr, nil))

		productUseCaseMock := mock.NewMockProductUseCase(ctrl)
		clientUseCaseMock := mock.NewMockClientUseCase(ctrl)
		orderUseCaseMock := mock.NewMockOrderUseCase(ctrl)

		productUseCaseMock.EXPECT().GetById(1).Return(nil, expectedErr).Times(1)

		c := NewOrderController(orderUseCaseMock, productUseCaseMock, clientUseCaseMock, loggerMock)
		c.CreateOrder(w, r)

		jsonResponse, _ := json.Marshal(Response{
			Error: "Error finding product",
			Data:  nil,
		})

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Equal(t, string(jsonResponse), string(w.Body.Bytes()))
	})

	t.Run("product not found error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		body := dto.OrderDto{
			ClientId: nil,
			Products: []struct {
				ID int `json:"id"`
			}{
				{
					ID: 1,
				},
			},
		}

		jsonBody, _ := json.Marshal(body)
		bodyReader := bytes.NewReader(jsonBody)

		r, _ := http.NewRequest("POST", "/clients", bodyReader)
		w := httptest.NewRecorder()

		loggerMock := slog.New(slog.NewTextHandler(os.Stderr, nil))

		productUseCaseMock := mock.NewMockProductUseCase(ctrl)
		clientUseCaseMock := mock.NewMockClientUseCase(ctrl)
		orderUseCaseMock := mock.NewMockOrderUseCase(ctrl)

		productUseCaseMock.EXPECT().GetById(1).Return(nil, nil).Times(1)

		c := NewOrderController(orderUseCaseMock, productUseCaseMock, clientUseCaseMock, loggerMock)
		c.CreateOrder(w, r)

		jsonResponse, _ := json.Marshal(Response{
			Error: "Product of id 1 not found",
			Data:  nil,
		})

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Equal(t, string(jsonResponse), string(w.Body.Bytes()))
	})

	t.Run("client not found error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		product := entity.Product{
			ID:          1,
			NameProduct: "Test product",
			Description: "Test product",
			Price:       123.45,
			CategoryId:  1,
			Active:      true,
		}

		client := entity.Client{
			ID:    1,
			Cpf:   12345678900,
			Name:  "Test Client",
			Email: "testclient@example.com",
		}

		body := dto.OrderDto{
			ClientId: &client.ID,
			Products: []struct {
				ID int `json:"id"`
			}{
				{
					ID: 1,
				},
			},
		}

		jsonBody, _ := json.Marshal(body)
		bodyReader := bytes.NewReader(jsonBody)

		r, _ := http.NewRequest("POST", "/clients", bodyReader)
		w := httptest.NewRecorder()

		loggerMock := slog.New(slog.NewTextHandler(os.Stderr, nil))

		productUseCaseMock := mock.NewMockProductUseCase(ctrl)
		clientUseCaseMock := mock.NewMockClientUseCase(ctrl)
		orderUseCaseMock := mock.NewMockOrderUseCase(ctrl)

		getProduct := productUseCaseMock.EXPECT().GetById(1).Return(&product, nil).Times(1)
		clientUseCaseMock.EXPECT().GetClientById(body.ClientId).Return(nil, nil).Times(1).After(getProduct)

		c := NewOrderController(orderUseCaseMock, productUseCaseMock, clientUseCaseMock, loggerMock)
		c.CreateOrder(w, r)

		jsonResponse, _ := json.Marshal(Response{
			Error: "Client not found",
			Data:  nil,
		})

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Equal(t, string(jsonResponse), string(w.Body.Bytes()))
	})

	t.Run("error creating order", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		expectedErr := errors.New("error connecting to database")

		product := entity.Product{
			ID:          1,
			NameProduct: "Test product",
			Description: "Test product",
			Price:       123.45,
			CategoryId:  1,
			Active:      true,
		}

		client := entity.Client{
			ID:    1,
			Cpf:   12345678900,
			Name:  "Test Client",
			Email: "testclient@example.com",
		}

		body := dto.OrderDto{
			ClientId: &client.ID,
			Products: []struct {
				ID int `json:"id"`
			}{
				{
					ID: 1,
				},
			},
		}

		products := []*entity.Product{&product}

		jsonBody, _ := json.Marshal(body)
		bodyReader := bytes.NewReader(jsonBody)

		r, _ := http.NewRequest("POST", "/clients", bodyReader)
		w := httptest.NewRecorder()

		loggerMock := slog.New(slog.NewTextHandler(os.Stderr, nil))

		productUseCaseMock := mock.NewMockProductUseCase(ctrl)
		clientUseCaseMock := mock.NewMockClientUseCase(ctrl)
		orderUseCaseMock := mock.NewMockOrderUseCase(ctrl)

		getProduct := productUseCaseMock.EXPECT().GetById(1).Return(&product, nil).Times(1)
		getClient := clientUseCaseMock.EXPECT().GetClientById(body.ClientId).Return(&client, nil).Times(1).After(getProduct)
		orderUseCaseMock.EXPECT().Create(body.ConvertToEntity(), products).Return(nil, expectedErr).Times(1).After(getClient)

		c := NewOrderController(orderUseCaseMock, productUseCaseMock, clientUseCaseMock, loggerMock)
		c.CreateOrder(w, r)

		jsonResponse, _ := json.Marshal(Response{
			Error: "Error creating order",
			Data:  nil,
		})

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, string(jsonResponse), string(w.Body.Bytes()))
	})
}
