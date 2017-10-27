package main

import (
	"context"
	"io"
	"io/ioutil"

	"github.com/segmentio/kafka-go"
)

// KafkaIntegration holds open connection to Kafka writer and allows
// caller to produce messages to Kafka stream.
type KafkaIntegration struct {
	producer *kafka.Writer

	Brokers []string
	Topic   string
}

// Init initializes the Kafka writer with brokers and topic
func (k *KafkaIntegration) Init() error {
	k.producer = kafka.NewWriter(kafka.WriterConfig{
		Brokers:  k.Brokers,
		Topic:    k.Topic,
		Balancer: &kafka.LeastBytes{},
	})

	return nil
}

// Process writes data to the Kafka stream
func (k *KafkaIntegration) Process(r io.ReadCloser) error {
	defer r.Close()

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	return k.producer.WriteMessages(context.Background(), kafka.Message{
		Value: b,
	})
}
