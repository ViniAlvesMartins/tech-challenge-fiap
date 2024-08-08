package sns

import (
	"context"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-common/sns"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/entity"
)

type Service struct {
	sns *sns.Service
}

func NewService(sns *sns.Service) *Service {
	return &Service{
		sns: sns,
	}
}

func (s *Service) SendMessage(ctx context.Context, message entity.Order) error {
	return s.sns.Publish(ctx, message)
}
