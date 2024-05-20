package service

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type SqsService struct{}

func NewSqsService() *SqsService { return &SqsService{} }

func NewConnection() (*sqs.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
	)

	if err != nil {
		return nil, err
	}

	client := sqs.NewFromConfig(cfg)

	return client, nil
}

func (s *SqsService) ReceiveMessage(queueURL string) (*types.Message, error) {
	client, _ := NewConnection()

	out, err := client.ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
		QueueUrl:            &queueURL,
		MaxNumberOfMessages: 1,
		WaitTimeSeconds:     20,
	})

	if err != nil {
		return nil, err
	}

	if len(out.Messages) >= 1 {
		return &out.Messages[0], nil
	}

	return nil, nil
}

func (s *SqsService) DeleteMessage(queueURL string, receiptHandle string) error {
	client, _ := NewConnection()

	if _, err := client.DeleteMessage(context.TODO(), &sqs.DeleteMessageInput{
		QueueUrl:      &queueURL,
		ReceiptHandle: &receiptHandle,
	}); err != nil {
		return err
	}

	return nil
}

func (s *SqsService) SendMessage(queueURL string, message string, messageGroupId string) error {
	client, _ := NewConnection()

	if _, err := client.SendMessage(context.TODO(), &sqs.SendMessageInput{
		QueueUrl:               &queueURL,
		MessageBody:            &message,
		MessageGroupId:         &messageGroupId,
		MessageDeduplicationId: &message,
	}); err != nil {
		return err
	}

	return nil
}
