package entity

type ClientEntity struct {
	Id    string `json:"id" gorm:"primaryKey"`
	Cpf   int32  `json:"cpf"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
