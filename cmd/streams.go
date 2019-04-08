package main

import (
	"fmt"
	"os"
	"os/signal"
	"streams/internal/kafka/consumer"
	"streams/internal/parser"
	"sync"
	"syscall"
)

func main() {
	//TODO Replace with context

	r := consumer.CreateConsumer()
	var wg sync.WaitGroup
	wg.Add(2)
	parserQueue := make(chan []byte)

	go consumer.Subscribe(&wg, r, parserQueue)
	go parser.Parse(&wg, parserQueue, 5)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, os.Interrupt, syscall.SIGTERM)

	<-sigterm // Await a sigterm signal before safely closing the consumer

	fmt.Println("Closing consumer group")
	_ = r.Close()
	wg.Wait()
	fmt.Println("Bye")
}
