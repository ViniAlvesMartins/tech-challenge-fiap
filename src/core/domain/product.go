package domain

type category string

//const (
//	snack     category = "lanche"
//	accompany category = "acompanhamento"
//	drink     category = "bebida"
//	dessert   category = "sobremesa"
//)

type Product struct {
	ID          int64   `json:"id" gorm:"primaryKey;autoIncrement"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Category    string  `json:"category" gorm:"index"`
}
