package core

type ClientDomainInterface interface {
	GetId() string
	GetCpf() int32
	GetName() string
	GetEmail() string

	SetId(string)
}

func NewClientDomain(
	id string, cpf int32, name string,
	email string,
) ClientDomainInterface {
	return &ClientDomain{
		Id:    id,
		Cpf:   cpf,
		Name:  name,
		Email: email,
	}
}