package converter

import (
	"fiappos/ViniAlvesMartins/tech-challenge-fiap/src/adapter/outbound/repository/entity"
	core "fiappos/ViniAlvesMartins/tech-challenge-fiap/src/core/domain"
)

func ConvertDomainToEntity(
	domain core.ClientDomainInterface,
) *entity.ClientEntity {
	return &entity.ClientEntity{
		Id:    domain.GetId(),
		Cpf:   domain.GetCpf(),
		Name:  domain.GetName(),
		Email: domain.GetEmail(),
	}
}
