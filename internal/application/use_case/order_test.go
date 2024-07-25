package use_case

import (
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

		order := entity.Order{
			ID:          1,
			ClientId:    nil,
			StatusOrder: enum.OrderStatusAwaitingPayment,
			Amount:      500.50,
			Products: []*entity.Product{
				{
					ID:          1,
					NameProduct: "Test product 1",
					Description: "Test product 1",
					Price:       250.25,
					CategoryId:  1,
					Active:      true,
				},
				{
					ID:          2,
					NameProduct: "Test product 2",
					Description: "Test product 2",
					Price:       250.25,
					CategoryId:  1,
					Active:      true,
				},
			},
		}

		repo := mock.NewMockOrderRepository(ctrl)
		repo.EXPECT().Create(order).Return(order, nil).Times(1)

		orderUseCase := NewOrderUseCase(repo)
		newOrder, err := orderUseCase.Create(order)

		assert.Equal(t, order, *newOrder)
		assert.Nil(t, err)
	})

	t.Run("error creating order", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		expectedErr := errors.New("error connecting to database")

		order := entity.Order{
			ID:          1,
			ClientId:    nil,
			StatusOrder: enum.OrderStatusAwaitingPayment,
			Amount:      500.50,
			Products: []*entity.Product{
				{
					ID:          1,
					NameProduct: "Test product 1",
					Description: "Test product 1",
					Price:       250.25,
					CategoryId:  1,
					Active:      true,
				},
				{
					ID:          2,
					NameProduct: "Test product 2",
					Description: "Test product 2",
					Price:       250.25,
					CategoryId:  1,
					Active:      true,
				},
			},
		}

		repo := mock.NewMockOrderRepository(ctrl)
		repo.EXPECT().Create(order).Return(entity.Order{}, expectedErr).Times(1)

		orderUseCase := NewOrderUseCase(repo)
		newOrder, err := orderUseCase.Create(order)

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
			StatusOrder: enum.OrderStatusAwaitingPayment,
			Amount:      500.50,
			Products: []*entity.Product{
				{
					ID:          1,
					NameProduct: "Test product 1",
					Description: "Test product 1",
					Price:       250.25,
					CategoryId:  1,
					Active:      true,
				},
				{
					ID:          2,
					NameProduct: "Test product 2",
					Description: "Test product 2",
					Price:       250.25,
					CategoryId:  1,
					Active:      true,
				},
			},
		}

		repo := mock.NewMockOrderRepository(ctrl)
		repo.EXPECT().GetById(1).Return(&order, nil).Times(1)

		useCase := NewOrderUseCase(repo)
		o, err := useCase.GetById(1)

		assert.Nil(t, err)
		assert.Equal(t, order, *o)
	})

	t.Run("error getting order by id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		expectedErr := errors.New("error connecting to database")

		repo := mock.NewMockOrderRepository(ctrl)
		repo.EXPECT().GetById(1).Return(nil, expectedErr).Times(1)

		useCase := NewOrderUseCase(repo)
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
				StatusOrder: enum.OrderStatusAwaitingPayment,
				Amount:      500.50,
				Products: []*entity.Product{
					{
						ID:          1,
						NameProduct: "Test product 1",
						Description: "Test product 1",
						Price:       250.25,
						CategoryId:  1,
						Active:      true,
					},
					{
						ID:          2,
						NameProduct: "Test product 2",
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
				StatusOrder: enum.OrderStatusAwaitingPayment,
				Amount:      100.50,
				Products: []*entity.Product{
					{
						ID:          1,
						NameProduct: "Test product 3",
						Description: "Test product 3",
						Price:       50.25,
						CategoryId:  1,
						Active:      true,
					},
					{
						ID:          2,
						NameProduct: "Test product 4",
						Description: "Test product 4",
						Price:       550.25,
						CategoryId:  1,
						Active:      true,
					},
				},
			},
		}

		repo := mock.NewMockOrderRepository(ctrl)
		repo.EXPECT().GetAll().Return(orders, nil).Times(1)

		useCase := NewOrderUseCase(repo)
		o, err := useCase.GetAll()

		assert.Nil(t, err)
		assert.Equal(t, orders, *o)
	})

	t.Run("error getting all orders", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		expectedErr := errors.New("error connecting to database")

		repo := mock.NewMockOrderRepository(ctrl)
		repo.EXPECT().GetAll().Return(nil, expectedErr).Times(1)

		useCase := NewOrderUseCase(repo)
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

		useCase := NewOrderUseCase(repo)
		err := useCase.UpdateStatusById(1, enum.OrderStatusPreparing)

		assert.Nil(t, err)
	})

	t.Run("error updating by status id successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		expectedErr := errors.New("error connecting to database")

		repo := mock.NewMockOrderRepository(ctrl)
		repo.EXPECT().UpdateStatusById(1, enum.OrderStatusPreparing).Times(1).Return(expectedErr)

		useCase := NewOrderUseCase(repo)
		err := useCase.UpdateStatusById(1, enum.OrderStatusPreparing)

		assert.Error(t, expectedErr, err)
	})
}
