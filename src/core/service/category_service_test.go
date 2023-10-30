package service

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/port/mock"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
	"log/slog"
	"os"
	"testing"
)

func TestCatgoryService(t *testing.T) {
	t.Run("get category by id successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockCategoryRepository(ctrl)

		expected := &entity.Category{
			ID:   1,
			Name: "Bebida",
		}

		service := NewCategoryService(repo, slog.New(slog.NewTextHandler(os.Stderr, nil)))
		repo.EXPECT().GetById(expected.ID).Times(1).Return(expected, nil)

		category, err := service.GetById(expected.ID)

		assert.Equal(t, expected, category)
		assert.Nil(t, err)
	})
}
