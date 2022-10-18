package pkg

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type IProducer interface {
	ProduceMessage(topic string, message string) error
}

type Producer struct {
	Writer *kafka.Producer
}

func NewProducer() *Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	CheckError(err)
	return &Producer{
		Writer: p,
	}
}

func (p *Producer) ProduceMessage(topic string, message string) error {
	err := p.Writer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)
	return err
}
