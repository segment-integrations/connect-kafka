package main

import (
	_ "net/http/pprof"

	"github.com/segmentio/connect"
)

const (
	Version = "0.0.1-beta"
)

var usage = `
Usage:
  connect-kafka
    --topic=<topic>
    --broker=<url>...
    [--trusted-cert=<path> --client-cert=<path> --client-cert-key=<path>]
  connect-kafka -h | --help
  connect-kafka --version

Options:
  -h --help                   Show this screen
  --version                   Show version
  --topic=<topic>             Kafka topic name
  --broker=<url>              Kafka broker URL
`

func main() {
	connect.Run(&KafkaIntegration{})
}
