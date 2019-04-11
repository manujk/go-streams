package main

import (
	"fmt"
	"os"
	"os/signal"
	"streams/internal/kafka/consumer"
	"streams/internal/kafka/producer"
	"streams/internal/parser"
	"streams/internal/utils"
	"sync"
	"syscall"
)

func main() {

	config := utils.LoadConfig()

	//TODO Replace with context

	r := consumer.CreateConsumer(config.Kafka.Brokers, config.Kafka.ConsumerTopic, config.Kafka.Group)
	w, _ := producer.Connect(config.Kafka.Brokers)

	var wg sync.WaitGroup
	parserQueue := make(chan []byte)
	sinkQueue := make(chan string)

	concurrency := 1
	wg.Add(concurrency)
	for i := 1; i <= concurrency; i++ {
		go func() {
			defer wg.Done() //no closure value needed as we will send the same value to each goroutine
			consumer.Subscribe(r, parserQueue)
		}()
	}

	concurrency = 2
	wg.Add(concurrency)
	for i := 1; i <= concurrency; i++ {
		go func() {
			defer wg.Done()
			parser.Parse(parserQueue, sinkQueue)
		}()
	}
	concurrency = 5
	wg.Add(concurrency)
	for i := 1; i <= concurrency; i++ {
		go func() {
			defer wg.Done()
			producer.Publish(w, config.Kafka.ProducerTopic, sinkQueue)
		}()
	}

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, os.Interrupt, syscall.SIGTERM)

	<-sigterm // Await a sigterm signal before safely closing the consumer

	fmt.Println("Closing consumer group")
	_ = r.Close()
	_ = w.Close()
	close(parserQueue)
	close(sinkQueue)
	wg.Wait()

	fmt.Println("Bye")
}
