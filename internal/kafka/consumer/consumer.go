package consumer

import (
	"context"
	"github.com/segmentio/kafka-go"
	"math/rand"
	"time"
)

func CreateConsumer(brokers []string, topic string, group string) *kafka.Reader {
	rand.Seed(time.Now().UnixNano())
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		//GroupID:        utils.RandSeq(10),
		GroupID:  group,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB

		/*TODO: Commit after the data has been processed completely{
			****Currently for location based campaign it will work
			****but need to fix for topup and topping
		}
		struct Message{
			m kafkaMessage
			ctx Context
			data parsedData
		}

		*/
		CommitInterval: 3 * time.Second, // 3 Seconds
	})

	_ = r.SetOffset(-1)

	return r
}

func Subscribe(r *kafka.Reader, out chan []byte) {

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		out <- m.Value
	}
}
