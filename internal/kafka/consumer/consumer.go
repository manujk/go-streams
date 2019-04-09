package consumer

import (
	"context"
	"github.com/segmentio/kafka-go"
	"math/rand"
	"sync"
	"time"
)

func CreateConsumer(brokers []string, topic string, group string) *kafka.Reader {
	rand.Seed(time.Now().UnixNano())
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		//GroupID:        utils.RandSeq(10),
		GroupID:        group,
		Topic:          topic,
		MinBytes:       10e3,            // 10KB
		MaxBytes:       10e6,            // 10MB
		CommitInterval: 3 * time.Second, // 3 Seconds
	})

	_ = r.SetOffset(-1)

	return r

}

func Subscribe(wg *sync.WaitGroup, r *kafka.Reader, ch chan []byte) {

	defer func(wg *sync.WaitGroup, ch chan []byte) {
		close(ch)
		wg.Done()
	}(wg, ch)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		ch <- m.Value
	}
}
