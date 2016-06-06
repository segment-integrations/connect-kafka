package api

import (
	"io/ioutil"
	"net/http"

	"github.com/Shopify/sarama"
	"github.com/gohttp/app"
	"github.com/gohttp/logger"
	"github.com/gohttp/response"
	"github.com/segmentio/kit/log"
)

// Api structure
type Server struct {
	server http.Handler
	*app.App
	kafkaTopic    string
	kafkaProducer sarama.SyncProducer
}

// New creates a new Server, initializes it and returns it.
func New(topic string, producer sarama.SyncProducer) *Server {
	api := &Server{App: app.New(), kafkaTopic: topic, kafkaProducer: producer}
	api.Use(logger.New())
	api.Post("/listen", api.requestHandler)
	return api
}

// requestHandler is an HTTP Handler function that listens on /set
// endpoint. Uses the `SetMessage` struct in the `message` package
// to unmarshal the message.
func (s *Server) requestHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}

	log.Debugf("Message: %v", string(b))

	_, _, err = s.kafkaProducer.SendMessage(&sarama.ProducerMessage{
		Topic: s.kafkaTopic,
		Value: sarama.ByteEncoder(b),
	})
	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}

	// Return success to user.
	response.OK(w)
}
