package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"

	"github.com/Shopify/sarama"
	log "github.com/Sirupsen/logrus"
	"github.com/segment-integrations/connect-kafka/internal/kafka"
	"github.com/tj/docopt"
)

type KafkaIntegration struct {
	topic    string
	producer sarama.SyncProducer
}

func (k *KafkaIntegration) newTLSFromConfig(m map[string]interface{}) *tls.Config {
	trustedCertPath, _ := m["--trusted-cert"].(string)
	clientCertPath, _ := m["--client-cert"].(string)
	clientCertKeyPath, _ := m["--client-cert-key"].(string)

	if trustedCertPath == "" && clientCertPath == "" && clientCertKeyPath == "" {
		return nil
	}

	trustedCertBytes, err := ioutil.ReadFile(trustedCertPath)
	if err != nil {
		log.Fatal(err)
	}

	clientCertBytes, err := ioutil.ReadFile(clientCertPath)
	if err != nil {
		log.Fatal(err)
	}

	clientCertKeyBytes, err := ioutil.ReadFile(clientCertKeyPath)
	if err != nil {
		log.Fatal(err)
	}

	cert, err := tls.X509KeyPair(clientCertBytes, clientCertKeyBytes)
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(trustedCertBytes)

	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
		RootCAs:            certPool,
	}
	tlsConfig.BuildNameToCertificate()

	return tlsConfig
}

func (k *KafkaIntegration) Init() error {
	m, err := docopt.Parse(usage, nil, true, Version, false)
	if err != nil {
		return err
	}

	kafkaConfig := &kafka.Config{BrokerAddresses: m["--broker"].([]string)}
	kafkaConfig.TLSConfig = k.newTLSFromConfig(m)

	producer, err := kafka.NewProducer(kafkaConfig)
	if err != nil {
		return err
	}

	k.producer = producer
	k.topic = m["--topic"].(string)

	return nil
}

func (k *KafkaIntegration) Process(r io.ReadCloser) error {
	defer r.Close()
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	_, _, err = k.producer.SendMessage(&sarama.ProducerMessage{
		Topic: k.topic,
		Value: sarama.ByteEncoder(b),
	})

	return err
}
