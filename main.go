package main

import (
	_ "net/http/pprof"

	"github.com/Sirupsen/logrus"
	"github.com/segmentio/connect"
	"github.com/tj/docopt"
)

const version = "0.0.1-beta"

var usage = `
Usage:
  connect-kafka
    --topic=<topic>
    --broker=<url>...
  connect-kafka -h | --help
  connect-kafka --version

Options:
  -h --help                   Show this screen
  --version                   Show version
  --topic=<topic>             Kafka topic name
  --broker=<url>              Kafka broker URL
`

func main() {
	m, err := docopt.Parse(usage, nil, true, version, false)
	if err != nil {
		logrus.Fatalf("parse flags error: %v", err)
	}

	var (
		brokers = m["--broker"].([]string)
		topic   = m["--topic"].(string)
	)

	connect.Run(&KafkaIntegration{
		Brokers: brokers,
		Topic:   topic,
	})
}
