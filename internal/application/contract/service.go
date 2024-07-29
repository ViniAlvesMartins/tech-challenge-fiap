//go:generate mockgen -destination=mock/service.go -source=service.go -package=mock
package contract

import (
	"context"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/entity"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type QueueService interface {
	ReceiveMessage(queueUrl string) (*types.Message, error)
	DeleteMessage(queueURL string, receiptHandle string) error
	SendMessage(queueUrl string, message string, messageGroupId string) error
}

type SnsService interface {
	SendMessage(ctx context.Context, message entity.Order) error
}
