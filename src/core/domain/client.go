package domain

type Client struct {
	Id    int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Cpf   int    `json:"cpf"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
