package application

import (
	"strings"
	"testing"

	kafka "github.com/segmentio/kafka-go"
)

type commentMockService struct {
	update map[string]error
}

func NewMockKafkaReader() *kafka.Reader {
	brokers := strings.Split("127.0.0.1", ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  "0",
		Topic:    "my-topic",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}
func NewMockCommentService() *commentMockService {
	return &commentMockService{
		update: map[string]error{},
	}
}

func (s *commentMockService) UpdateComment(msg kafka.Message) error {
	return s.update[string(msg.Value)]
}

func TestConsuming(t *testing.T) {
	s := NewMockCommentService()
	r := NewMockKafkaReader()
	c := NewCommentConsume(s)

	c.Consuming(r)
}
