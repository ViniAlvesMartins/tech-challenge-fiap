package core

type ClientDomain struct {
	Id    string
	Cpf   int32
	Name  string
	Email string
}

func (ud *ClientDomain) GetId() string {
	return ud.Id
}

func (ud *ClientDomain) SetId(Id string) {
	ud.Id = Id
}

func (ud *ClientDomain) GetCpf() int32 {
	return ud.Cpf
}

func (ud *ClientDomain) SetCpf(Cpf int32) {
	ud.Cpf = Cpf
}

func (ud *ClientDomain) GetName() string {
	return ud.Name
}

func (ud *ClientDomain) SetName(Name string) {
	ud.Name = Name
}

func (ud *ClientDomain) GetEmail() string {
	return ud.Email
}

func (ud *ClientDomain) SetEmail(Email string) {
	ud.Email = Email
}
