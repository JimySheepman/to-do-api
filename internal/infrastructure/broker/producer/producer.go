package producer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	kafka "github.com/segmentio/kafka-go"
)

const (
	KAFKA_URL = "127.0.0.1"
	TOPIC     = "my-topic"
)

func NewKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func Send(key string, value interface{}) error {
	writer := NewKafkaWriter(KAFKA_URL, TOPIC)
	defer writer.Close()

	buff := new(bytes.Buffer)
	json.NewEncoder(buff).Encode(value)

	fmt.Println("start producing sending ... !!")

	msg := kafka.Message{
		Key:   []byte(key),
		Value: buff.Bytes(),
	}

	err := writer.WriteMessages(context.Background(), msg)
	if err != nil {
		return err
	} else {
		fmt.Println("produced -> ", key)
	}

	defer fmt.Println("close producing sending ... !!")

	return nil
}
