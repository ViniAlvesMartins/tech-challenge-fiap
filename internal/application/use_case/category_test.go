package use_case

import (
	"errors"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/application/contract/mock"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/entity"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"os"
	"testing"
)

func TestNewCategoryUseCase(t *testing.T) {
	t.Run("get by id successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

		category := &entity.Category{
			ID:   1,
			Name: "Test category",
		}

		repo := mock.NewMockCategoryRepository(ctrl)
		repo.EXPECT().GetById(1).Return(category, nil).Times(1)

		categoryUseCase := NewCategoryUseCase(repo, logger)
		c, err := categoryUseCase.GetById(1)

		assert.Equal(t, *category, *c)
		assert.Nil(t, err)
	})

	t.Run("database error getting category by id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
		expectedErr := errors.New("database error")

		repo := mock.NewMockCategoryRepository(ctrl)
		repo.EXPECT().GetById(1).Return(nil, expectedErr).Times(1)

		categoryUseCase := NewCategoryUseCase(repo, logger)
		c, err := categoryUseCase.GetById(1)

		assert.ErrorIs(t, expectedErr, err)
		assert.Nil(t, c)
	})
}
