package sqs

import (
	"encoding/json"
	"fmt"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/enum"
	"log"
	"log/slog"
	"time"
)

type SqsConsumer struct {
	sqsService   contract.QueueService
	orderUseCase contract.OrderUseCase
	logger       *slog.Logger
}

func NewSqsConsumer(queueService contract.QueueService,
	orderUseCase contract.OrderUseCase,
	logger *slog.Logger) *SqsConsumer {
	return &SqsConsumer{
		sqsService:   queueService,
		orderUseCase: orderUseCase,
		logger:       logger,
	}
}

func (s *SqsConsumer) Run() error {

	fmt.Println("teste aqui")

	queueUrl := "https://sqs.us-east-1.amazonaws.com/682279319757/from_payment_production_order_queue"

	for {

		result, err := s.sqsService.ReceiveMessage(queueUrl)

		if err != nil {
			log.Printf("Failed to fetch sqs message %v", err)
			continue
		}

		if result == nil {
			log.Printf("Failed result sqs %v", result)
			continue
		} else {

			sqsMessage := &SqsMessage{}

			json.Unmarshal([]byte(*result.Body), &sqsMessage)

			sqsMessageReturn := &SqsMessageReturn{}

			log.Println(sqsMessage.Message)

			json.Unmarshal([]byte(sqsMessage.Message), &sqsMessageReturn)

			status := ""

			if sqsMessageReturn.Status == "CONFIRMED" {
				status = string(enum.PREPARING)
			} else if sqsMessageReturn.Status == "FINISHED" {
				status = string(enum.FINISHED)
			}

			log.Println(sqsMessageReturn)

			err := s.orderUseCase.UpdateStatusById(sqsMessageReturn.OrderId, enum.StatusOrder(status))

			if err != nil {
				s.logger.Error("error updating status by id", slog.Any("error", err.Error()))
			} else {
				log.Printf(*result.ReceiptHandle)
				s.sqsService.DeleteMessage(queueUrl, *result.ReceiptHandle)

				if sqsMessageReturn.Status == "CONFIRMED" {

					js, _ := json.Marshal(&SqsMessageReturn{
						OrderId: sqsMessageReturn.OrderId,
						Status:  string(enum.PREPARING),
					})

					s.sqsService.SendMessage("https://sqs.us-east-1.amazonaws.com/682279319757/to_production_order_queue.fifo",
						string(js),
						fmt.Sprint(sqsMessageReturn.OrderId))
				}

			}

		}

		time.Sleep(5 * time.Second)
	}
}

type SqsMessage struct {
	Type             string
	MessageId        string
	TopicArn         string
	Message          string
	Timestamp        string
	SignatureVersion string
	Signature        string
	SigningCertURL   string
	UnsubscribeURL   string
}

type SqsMessageReturn struct {
	OrderId int
	Status  string
}
