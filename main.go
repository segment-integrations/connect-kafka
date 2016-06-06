package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net"
	"net/http"

	_ "net/http/pprof"

	log "github.com/Sirupsen/logrus"
	"github.com/segment-integrations/connect-kafka/internal/api"
	"github.com/segment-integrations/connect-kafka/internal/kafka"
	"github.com/tj/docopt"
)

const (
	Version = "0.0.1-beta"
)

var usage = `
Usage:
  connect-kafka
    [--debug]
    --topic=<topic>
    --broker=<url>...
    [--listen=<addr>]
    [--trusted-cert=<path> --client-cert=<path> --client-cert-key=<path>]
  connect-kafka -h | --help
  connect-kafka --version

Options:
  -h --help                   Show this screen
  --version                   Show version
  --topic=<topic>             Kafka topic name
  --listen=<addr>             Address to listen on [default: localhost:3000]
  --broker=<url>              Kafka broker URL
`

func newTLSFromConfig(m map[string]interface{}) *tls.Config {
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

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	m, err := docopt.Parse(usage, nil, true, Version, false)
	if err != nil {
		log.Error(err)
		return
	}

	kafkaConfig := &kafka.Config{BrokerAddresses: m["--broker"].([]string)}
	kafkaConfig.TLSConfig = newTLSFromConfig(m)

	producer, err := kafka.NewProducer(kafkaConfig)
	if err != nil {
		log.Error(err)
		return
	}

	// Server configuration
	api := api.New(m["--topic"].(string), producer)

	// Create listener
	listener, err := net.Listen("tcp", m["--listen"].(string))
	if err != nil {
		log.Error(err)
		return
	}

	if m["--debug"].(bool) {
		log.SetLevel(log.DebugLevel)
	}

	// Run listen and serve
	log.Infof("Server started at %v", listener.Addr())
	log.Fatal(http.Serve(listener, api))
}
