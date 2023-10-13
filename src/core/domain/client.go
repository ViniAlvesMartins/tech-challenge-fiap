package domain

type Client struct {
	Id    string `json:"id"`
	Cpf   int32  `json:"cpf"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
