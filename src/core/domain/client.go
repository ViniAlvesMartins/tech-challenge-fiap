package domain

type Client struct {
	Id    int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Cpf   int32  `json:"cpf"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
