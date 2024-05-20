package use_case

import (
	"errors"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/application/contract/mock"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/entity"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"os"
	"testing"
)

func TestNewClientUseCase_Create(t *testing.T) {
	t.Run("create client successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

		client := entity.Client{
			ID:    1,
			Cpf:   12345678900,
			Name:  "Test Client",
			Email: "client@example.com",
		}

		repo := mock.NewMockClientRepository(ctrl)
		repo.EXPECT().Create(client).Return(client, nil).Times(1)

		clientUseCase := NewClientUseCase(repo, logger)

		newClient, err := clientUseCase.Create(client)

		assert.Equal(t, client, *newClient)
		assert.Nil(t, err)
	})

	t.Run("error saving to database", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
		expectedErr := errors.New("error connecting to database")

		client := entity.Client{
			ID:    1,
			Cpf:   12345678900,
			Name:  "Test Client",
			Email: "client@example.com",
		}

		repo := mock.NewMockClientRepository(ctrl)
		repo.EXPECT().Create(client).Return(entity.Client{}, expectedErr).Times(1)

		clientUseCase := NewClientUseCase(repo, logger)

		newClient, err := clientUseCase.Create(client)

		assert.Nil(t, newClient)
		assert.ErrorIs(t, expectedErr, err)
	})
}

func TestNewClientUseCase_GetClientByCpf(t *testing.T) {
	t.Run("get client by cpf successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

		client := entity.Client{
			ID:    1,
			Cpf:   12345678900,
			Name:  "Test Client",
			Email: "client@example.com",
		}

		repo := mock.NewMockClientRepository(ctrl)
		repo.EXPECT().GetClientByCpf(client.Cpf).Return(&client, nil).Times(1)

		clientUseCase := NewClientUseCase(repo, logger)

		newClient, err := clientUseCase.GetClientByCpf(client.Cpf)

		assert.Equal(t, client, *newClient)
		assert.Nil(t, err)
	})

	t.Run("error getting client by cpf", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
		expectedErr := errors.New("error connecting to database")

		client := entity.Client{
			ID:    1,
			Cpf:   12345678900,
			Name:  "Test Client",
			Email: "client@example.com",
		}

		repo := mock.NewMockClientRepository(ctrl)
		repo.EXPECT().GetClientByCpf(client.Cpf).Return(nil, expectedErr).Times(1)

		clientUseCase := NewClientUseCase(repo, logger)

		newClient, err := clientUseCase.GetClientByCpf(client.Cpf)

		assert.Nil(t, newClient)
		assert.ErrorIs(t, expectedErr, err)
	})
}

func TestNewClientUseCase_GetClientById(t *testing.T) {
	t.Run("get client by id successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

		client := entity.Client{
			ID:    1,
			Cpf:   12345678900,
			Name:  "Test Client",
			Email: "client@example.com",
		}

		repo := mock.NewMockClientRepository(ctrl)
		repo.EXPECT().GetClientById(&client.ID).Return(&client, nil).Times(1)

		clientUseCase := NewClientUseCase(repo, logger)

		newClient, err := clientUseCase.GetClientById(&client.ID)

		assert.Equal(t, client, *newClient)
		assert.Nil(t, err)
	})

	t.Run("error getting client by id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
		expectedErr := errors.New("error connecting to database")

		client := entity.Client{
			ID:    1,
			Cpf:   12345678900,
			Name:  "Test Client",
			Email: "client@example.com",
		}

		repo := mock.NewMockClientRepository(ctrl)
		repo.EXPECT().GetClientById(&client.ID).Return(nil, expectedErr).Times(1)

		clientUseCase := NewClientUseCase(repo, logger)

		newClient, err := clientUseCase.GetClientById(&client.ID)

		assert.Nil(t, newClient)
		assert.ErrorIs(t, expectedErr, err)
	})
}

func TestNewClientUseCase_GetAlreadyExists(t *testing.T) {
	t.Run("check if client already exists successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

		client := entity.Client{
			ID:    1,
			Cpf:   12345678900,
			Name:  "Test Client",
			Email: "client@example.com",
		}

		repo := mock.NewMockClientRepository(ctrl)
		repo.EXPECT().GetAlreadyExists(client.Cpf, client.Email).Return(&client, nil).Times(1)

		clientUseCase := NewClientUseCase(repo, logger)

		newClient, err := clientUseCase.GetAlreadyExists(client.Cpf, client.Email)

		assert.Equal(t, client, *newClient)
		assert.Nil(t, err)
	})

	t.Run("error checking if client already exists", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
		expectedErr := errors.New("error connecting to database")

		client := entity.Client{
			ID:    1,
			Cpf:   12345678900,
			Name:  "Test Client",
			Email: "client@example.com",
		}

		repo := mock.NewMockClientRepository(ctrl)
		repo.EXPECT().GetAlreadyExists(client.Cpf, client.Email).Return(nil, expectedErr).Times(1)

		clientUseCase := NewClientUseCase(repo, logger)

		newClient, err := clientUseCase.GetAlreadyExists(client.Cpf, client.Email)

		assert.Nil(t, newClient)
		assert.ErrorIs(t, expectedErr, err)
	})
}