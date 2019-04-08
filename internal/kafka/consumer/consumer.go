package consumer

import (
	"context"
	"github.com/segmentio/kafka-go"
	"sync"
)

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
