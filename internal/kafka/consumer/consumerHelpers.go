package consumer

import (
	"github.com/segmentio/kafka-go"
	"math/rand"
	"streams/internal/utils"
	"time"
)

func CreateConsumer() *kafka.Reader {
	rand.Seed(time.Now().UnixNano())
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{"192.168.1.142:9092", "192.168.1.143:9092", "192.168.1.144:9092"},
		GroupID:        utils.RandSeq(10),
		Topic:          "gotest",
		MinBytes:       10e3,            // 10KB
		MaxBytes:       10e6,            // 10MB
		CommitInterval: 3 * time.Second, // 3 Seconds
	})

	_ = r.SetOffset(-1)

	return r

}
