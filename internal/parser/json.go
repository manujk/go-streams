package parser

import (
	"encoding/json"
	"fmt"
	"github.com/gammazero/workerpool"
	"streams/internal/model"
	"sync"
)

func Parse(wg *sync.WaitGroup, msgs chan []byte, concurrency int) {
	parserPool := workerpool.New(concurrency)
	defer func() {
		parserPool.StopWait()
		wg.Done()
	}()

	for {
		select {
		case msg, ok := <-msgs:
			if !ok {
				fmt.Println("All work done")
				return
			}

			parserPool.Submit(func() {
				var topupJSON model.TopupSchema
				_ = json.Unmarshal(msg, &topupJSON)
				fmt.Println("Msisdn", topupJSON.MSISDN)
			})

		}

	}

}
