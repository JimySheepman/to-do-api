package main

import (
	"context"
	"fmt"

	"github.com/JimySheepman/to-do-api/internal/infrastructure/broker/producer"
	kafka "github.com/segmentio/kafka-go"
)

func main() {
	kafkaURL := "127.0.0.1"
	topic := "my-topic"
	writer := producer.NewKafkaWriter(kafkaURL, topic)
	defer writer.Close()
	fmt.Println("start producing ... !!")
	for i := 0; ; i++ {
		key := fmt.Sprintf("Key-%d", i)
		msg := kafka.Message{
			Key:   []byte(key),
			Value: []byte(key),
		}
		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("produced", key)
		}
		//time.Sleep(1 * time.Second)
	}
}
