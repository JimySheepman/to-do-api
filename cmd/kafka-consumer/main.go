package main

import (
	"context"
	"fmt"
	"log"

	"github.com/JimySheepman/to-do-api/internal/infrastructure/broker/consumer"
)

func main() {
	kafkaURL := "127.0.0.1"
	topic := "my-topic"
	groupID := "0"

	reader := consumer.GetKafkaReader(kafkaURL, topic, groupID)

	defer reader.Close()

	fmt.Println("start consuming ... !!")
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
