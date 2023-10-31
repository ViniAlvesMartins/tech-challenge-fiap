//go:generate mockgen -destination=mock/category.go -source=category.go -package=mock
package port

import "github.com/ViniAlvesMartins/tech-challenge-fiap/src/core/domain/entity"

type CategoryRepository interface {
	GetById(id int) (*entity.Category, error)
}

type CategoryService interface {
	GetById(id int) (*entity.Category, error)
}
