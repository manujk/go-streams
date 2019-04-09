package parser

import (
	"encoding/json"
	"fmt"
	"github.com/gammazero/workerpool"
	"streams/internal/model"
	"sync"
)

func Parse(wg *sync.WaitGroup, in chan []byte, out chan string, concurrency int) {
	parserPool := workerpool.New(concurrency)
	defer func() {
		close(out)
		parserPool.StopWait()
		wg.Done()
	}()

	for {
		select {
		case msg, ok := <-in:
			if !ok {
				fmt.Println("All work done")
				return
			}

			parserPool.Submit(func() {
				var topupJSON model.TopupSchema
				_ = json.Unmarshal(msg, &topupJSON)
				fmt.Println("Msisdn", topupJSON.MSISDN)
				jsonString, _ := json.Marshal(topupJSON)
				out <- string(jsonString)
			})

		}

	}

}
