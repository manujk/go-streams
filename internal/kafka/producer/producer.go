package producer

import (
	"context"
	"fmt"
	"github.com/gammazero/workerpool"
	"github.com/segmentio/kafka-go"
	"sync"
)

func Connect(brokers []string, topic string) *kafka.Writer {

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  brokers,
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	return w
}

func MessageBuilder(key []byte, value string) kafka.Message {
	return kafka.Message{
		Key:   key,
		Value: []byte(value),
	}
}

func Publish(wg *sync.WaitGroup, w *kafka.Writer, in chan string, concurrency int) {

	producerPool := workerpool.New(concurrency)
	defer func() {
		producerPool.StopWait()
		_ = w.Close()
		wg.Done()
	}()

	for {
		select {
		case msg, ok := <-in:
			if !ok {
				fmt.Println("All work done")
				return
			}

			producerPool.Submit(func() {

				kafkaMsg := MessageBuilder(nil, msg)
				_ = w.WriteMessages(context.Background(), kafkaMsg)

			})

		}

	}

}
