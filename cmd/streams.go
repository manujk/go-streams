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

	r := consumer.CreateConsumer(config.Kafka.Brokers, config.Kafka.Topic, config.Kafka.Group)
	w := producer.Connect(config.Kafka.Brokers, config.Kafka.Topic)

	var wg sync.WaitGroup
	wg.Add(3)
	parserQueue := make(chan []byte)
	sinkQueue := make(chan string)

	go consumer.Subscribe(&wg, r, parserQueue)
	go parser.Parse(&wg, parserQueue, sinkQueue, 5)
	go producer.Publish(&wg, w, sinkQueue, 5)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, os.Interrupt, syscall.SIGTERM)

	<-sigterm // Await a sigterm signal before safely closing the consumer

	fmt.Println("Closing consumer group")
	_ = r.Close()
	wg.Wait()
	fmt.Println("Bye")
}
