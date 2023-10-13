package core

type ClientDomain struct {
	id    string
	cpf   int32
	name  string
	email string
}

func (ud *ClientDomain) GetId() string {
	return ud.id
}

func (ud *ClientDomain) SetId(id string) {
	ud.id = id
}

func (ud *ClientDomain) GetCpf() int32 {
	return ud.cpf
}

func (ud *ClientDomain) SetCpf(cpf int32) {
	ud.cpf = cpf
}

func (ud *ClientDomain) GetName() string {
	return ud.name
}

func (ud *ClientDomain) SetName(name string) {
	ud.name = name
}

func (ud *ClientDomain) GetEmail() string {
	return ud.email
}

func (ud *ClientDomain) SetEmail(email string) {
	ud.email = email
}
