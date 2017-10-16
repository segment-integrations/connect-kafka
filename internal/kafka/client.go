package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/satori/go.uuid"
)

func NewProducer(c *Config) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Errors = true
	config.Producer.Return.Successes = true
	config.ClientID = uuid.NewV4().String()

	// TLS
	if c.TLSConfig != nil {
		config.Net.TLS.Config = c.TLSConfig
		config.Net.TLS.Enable = true
	}

	err := config.Validate()
	if err != nil {
		return nil, err
	}

	producer, err := sarama.NewSyncProducer(c.getBrokers(), config)
	if err != nil {
		return nil, err
	}

	return producer, nil
}
