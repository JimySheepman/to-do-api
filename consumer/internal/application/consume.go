package application

import (
	"context"
	"log"

	"github.com/JimySheepman/to-do-api/consumer/internal/service"
	kafka "github.com/segmentio/kafka-go"
)

type CommentConsume struct {
	commentService service.CommentService
}

func NewCommentConsume(reader *kafka.Reader, service service.CommentService) {

	consume := &CommentConsume{
		commentService: service,
	}
	consume.Consuming(reader)
}

func (c *CommentConsume) Consuming(reader *kafka.Reader) error {
	defer reader.Close()

	log.Println("start consuming ... !!")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			return err
		}

		err = c.commentService.UpdateComment(msg)

		if err != nil {
			return err
		}

		log.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
	}
}
