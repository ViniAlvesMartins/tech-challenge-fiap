package use_case

import (
	"context"
	"errors"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/application/contract/mock"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/enum"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderUseCase_Create(t *testing.T) {
	t.Run("create order successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		ctx := context.Background()

		order := entity.Order{
			ID:          1,
			ClientId:    nil,
			OrderStatus: enum.OrderStatusAwaitingPayment,
			Amount:      500.50,
			Products: []*entity.Product{
				{
					ID:          1,
					ProductName: "Test product 1",
					Description: "Test product 1",
					Price:       250.25,
					CategoryId:  1,
					Active:      true,
				},
				{
					ID:          2,
					ProductName: "Test product 2",
					Description: "Test product 2",
					Price:       250.25,
					CategoryId:  1,
					Active:      true,
				},
			},
		}

		repo := mock.NewMockOrderRepository(ctrl)
		repo.EXPECT().Create(order).Return(order, nil).Times(1)

		snsService := mock.NewMockSnsService(ctrl)
		snsService.EXPECT().SendMessage(ctx, order).Return(nil)

		productUseCase := mock.NewMockProductUseCase(ctrl)
		productUseCase.EXPECT().GetById(gomock.Any()).Return(order.Products[1], nil).Times(2)

		orderUseCase := NewOrderUseCase(repo, productUseCase, snsService)
		newOrder, err := orderUseCase.Create(ctx, order)

		assert.Equal(t, order, *newOrder)
		assert.Nil(t, err)
	})

	t.Run("error creating order", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		ctx := context.Background()

		expectedErr := errors.New("error connecting to database")

		order := entity.Order{
			ID:          1,
			ClientId:    nil,
			OrderStatus: enum.OrderStatusAwaitingPayment,
			Amount:      500.50,
			Products: []*entity.Product{
				{
					ID:          1,
					ProductName: "Test product 1",
					Description: "Test product 1",
					Price:       250.25,
					CategoryId:  1,
					Active:      true,
				},
				{
					ID:          2,
					ProductName: "Test product 2",
					Description: "Test product 2",
					Price:       250.25,
					CategoryId:  1,
					Active:      true,
				},
			},
		}

		snsService := mock.NewMockSnsService(ctrl)

		productUseCase := mock.NewMockProductUseCase(ctrl)
		productUseCase.EXPECT().GetById(gomock.Any()).Return(order.Products[1], nil).Times(2)

		repo := mock.NewMockOrderRepository(ctrl)
		repo.EXPECT().Create(order).Return(entity.Order{}, expectedErr).Times(1)

		orderUseCase := NewOrderUseCase(repo, productUseCase, snsService)
		newOrder, err := orderUseCase.Create(ctx, order)

		assert.Equal(t, expectedErr, err)
		assert.Nil(t, newOrder)
	})
}

func TestOrderUseCase_GetById(t *testing.T) {
	t.Run("get order by id successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		order := entity.Order{
			ID:          1,
			ClientId:    nil,
			OrderStatus: enum.OrderStatusAwaitingPayment,
			Amount:      500.50,
			Products: []*entity.Product{
				{
					ID:          1,
					ProductName: "Test product 1",
					Description: "Test product 1",
					Price:       250.25,
					CategoryId:  1,
					Active:      true,
				},
				{
					ID:          2,
					ProductName: "Test product 2",
					Description: "Test product 2",
					Price:       250.25,
					CategoryId:  1,
					Active:      true,
				},
			},
		}

		snsService := mock.NewMockSnsService(ctrl)

		productUseCase := mock.NewMockProductUseCase(ctrl)

		repo := mock.NewMockOrderRepository(ctrl)
		repo.EXPECT().GetById(1).Return(&order, nil).Times(1)

		useCase := NewOrderUseCase(repo, productUseCase, snsService)
		o, err := useCase.GetById(1)

		assert.Nil(t, err)
		assert.Equal(t, order, *o)
	})

	t.Run("error getting order by id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		expectedErr := errors.New("error connecting to database")

		productUseCase := mock.NewMockProductUseCase(ctrl)

		snsService := mock.NewMockSnsService(ctrl)

		repo := mock.NewMockOrderRepository(ctrl)
		repo.EXPECT().GetById(1).Return(nil, expectedErr).Times(1)

		useCase := NewOrderUseCase(repo, productUseCase, snsService)
		o, err := useCase.GetById(1)

		assert.Nil(t, o)
		assert.Error(t, expectedErr, err)
	})
}

func TestOrderUseCase_GetByAll(t *testing.T) {
	t.Run("get order all successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		orders := []entity.Order{
			{
				ID:          1,
				ClientId:    nil,
				OrderStatus: enum.OrderStatusAwaitingPayment,
				Amount:      500.50,
				Products: []*entity.Product{
					{
						ID:          1,
						ProductName: "Test product 1",
						Description: "Test product 1",
						Price:       250.25,
						CategoryId:  1,
						Active:      true,
					},
					{
						ID:          2,
						ProductName: "Test product 2",
						Description: "Test product 2",
						Price:       250.25,
						CategoryId:  1,
						Active:      true,
					},
				},
			},
			{
				ID:          2,
				ClientId:    nil,
				OrderStatus: enum.OrderStatusAwaitingPayment,
				Amount:      100.50,
				Products: []*entity.Product{
					{
						ID:          1,
						ProductName: "Test product 3",
						Description: "Test product 3",
						Price:       50.25,
						CategoryId:  1,
						Active:      true,
					},
					{
						ID:          2,
						ProductName: "Test product 4",
						Description: "Test product 4",
						Price:       550.25,
						CategoryId:  1,
						Active:      true,
					},
				},
			},
		}

		productUseCase := mock.NewMockProductUseCase(ctrl)

		snsService := mock.NewMockSnsService(ctrl)

		repo := mock.NewMockOrderRepository(ctrl)
		repo.EXPECT().GetAll().Return(orders, nil).Times(1)

		useCase := NewOrderUseCase(repo, productUseCase, snsService)
		o, err := useCase.GetAll()

		assert.Nil(t, err)
		assert.Equal(t, orders, *o)
	})

	t.Run("error getting all orders", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		expectedErr := errors.New("error connecting to database")

		productUseCase := mock.NewMockProductUseCase(ctrl)

		snsService := mock.NewMockSnsService(ctrl)

		repo := mock.NewMockOrderRepository(ctrl)
		repo.EXPECT().GetAll().Return(nil, expectedErr).Times(1)

		useCase := NewOrderUseCase(repo, productUseCase, snsService)
		o, err := useCase.GetAll()

		assert.Nil(t, o)
		assert.Error(t, expectedErr, err)
	})
}

func TestOrderUseCase_UpdateStatusById(t *testing.T) {
	t.Run("update by status id successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		repo := mock.NewMockOrderRepository(ctrl)
		repo.EXPECT().UpdateStatusById(1, enum.OrderStatusPreparing).Times(1).Return(nil)

		productUseCase := mock.NewMockProductUseCase(ctrl)

		snsService := mock.NewMockSnsService(ctrl)

		useCase := NewOrderUseCase(repo, productUseCase, snsService)
		err := useCase.UpdateStatusById(1, enum.OrderStatusPreparing)

		assert.Nil(t, err)
	})

	t.Run("error updating by status id successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		expectedErr := errors.New("error connecting to database")

		productUseCase := mock.NewMockProductUseCase(ctrl)

		repo := mock.NewMockOrderRepository(ctrl)
		repo.EXPECT().UpdateStatusById(1, enum.OrderStatusPreparing).Times(1).Return(expectedErr)

		snsService := mock.NewMockSnsService(ctrl)

		useCase := NewOrderUseCase(repo, productUseCase, snsService)
		err := useCase.UpdateStatusById(1, enum.OrderStatusPreparing)

		assert.Error(t, expectedErr, err)
	})
}
