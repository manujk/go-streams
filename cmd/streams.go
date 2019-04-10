package main

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"os"
	"os/signal"
	"streams/internal/kafka/producer"
	"streams/internal/utils"
	"sync"
	"syscall"
)

func main() {

	config := utils.LoadConfig()

	version, err := sarama.ParseKafkaVersion("2.1.1")
	if err != nil {
		panic(err)
	}
	consumerConfig := sarama.NewConfig()
	consumerConfig.Version = version
	consumer := Consumer{}

	client, err := sarama.NewConsumerGroup(config.Kafka.Brokers, config.Kafka.Group, consumerConfig)
	if err != nil {
		panic(err)
	}
	consumer.ready = make(chan bool, 0)

	go func() {
		for {
			err := client.Consume(context.Background(), config.Kafka.Consumer, &consumer)
			if err != nil {
				return
			}
		}
	}()

	<-consumer.ready
	fmt.Println("Kafka consumer up and running!...")

	w, err := producer.Connect(config.Kafka.Brokers)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	parserQueue := make(chan []byte)
	sinkQueue := make(chan string)

	//concurrency := 2
	//wg.Add(concurrency)
	//for i := 1; i <= concurrency; i++ {
	//	go func(wg *sync.WaitGroup, in chan []byte, out chan string) {
	//		defer wg.Done()
	//		parser.Parse(in, out)
	//	}(&wg, parserQueue, sinkQueue)
	//}
	//
	//concurrency = 5
	//wg.Add(concurrency)
	//for i := 1; i <= concurrency; i++ {
	//	go func(wg *sync.WaitGroup, p sarama.SyncProducer, topic string, in chan string) {
	//		defer wg.Done()
	//		producer.Publish(p, topic, in)
	//	}(&wg, w, config.Kafka.Producer, sinkQueue)
	//}

	//for j := range jobs {
	//	doWork(i, j)
	//}

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, os.Interrupt, syscall.SIGTERM)

	<-sigterm // Await a sigterm signal before safely closing the consumer

	fmt.Println("Closing consumer group")
	err = client.Close()
	if err != nil {
		fmt.Println(err)
	}
	close(parserQueue)
	close(sinkQueue)
	wg.Wait()
	err = w.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Bye")
}

type Consumer struct {
	ready chan bool
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
	for message := range claim.Messages() {
		fmt.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
		session.MarkMessage(message, "")
	}

	return nil
}
