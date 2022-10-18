package pkg

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type IConsumer interface {
	ConsumeMessage(topic string) error
}

type Consumer struct {
	Reader *kafka.Consumer
}

func NewConsumer() *Consumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	CheckError(err)
	return &Consumer{
		Reader: c,
	}
}

func (c *Consumer) ConsumeMessage(topic string) error {
	c.Reader.SubscribeTopics([]string{topic}, nil)
	for {
		msg, err := c.Reader.ReadMessage(-1)
		if err == nil {
			log.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}
	c.Reader.Close()
	return nil
}
