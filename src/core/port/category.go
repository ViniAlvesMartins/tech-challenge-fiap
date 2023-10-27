package port

import "github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/entity"

type CategoryRepository interface {
	GetById(id int) (entity.Category, error)
}

type CategoryService interface {
	GetCategoryById(id int) (entity.Category, error)
}
