package converter

import (
	"fiappos/ViniAlvesMartins/tech-challenge-fiap/src/adapter/outbound/repository/entity"
	core "fiappos/ViniAlvesMartins/tech-challenge-fiap/src/core/domain"
)

func ConvertEntityToDomain(
	entity entity.ClientEntity,
) core.ClientDomainInterface {
	domain := core.NewClientDomain(
		entity.Id,
		entity.Cpf,
		entity.Name,
		entity.Email,
	)
	return domain
}
