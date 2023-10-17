package dto

type ProductDto struct {
	ID          int     `json:"id" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float32 `json:"price" validate:"required"`
	Category    string  `json:"category" validate:"required"`
}
