package input

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/entities/enum"
)

type PaymentDto struct {
	Order int    `json:"order" validate:"required" error:"ID do pedido é obrigatorio"`
	Type  string `json:"type" validate:"required" error:"Tipo de pagamento é obrigatorio"`
}

func (p *PaymentDto) ConvertToEntity() entity.Payment {
	return entity.Payment{
		OrderID: p.Order,
		Type:    enum.PaymentType(p.Type),
	}
}
