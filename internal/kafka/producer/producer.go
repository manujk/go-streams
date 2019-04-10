package producer

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func Connect(brokers []string) (sarama.SyncProducer, error) {

	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	fmt.Println()

	prd, err := sarama.NewSyncProducer(brokers, config)
	return prd, err
}

func Publish(producer sarama.SyncProducer, topic string, in chan string) {

	//channel range doesnt stop when there’s no more elements to be read
	// unless the channel is closed
	count := 0
	for message := range in {
		count++
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(message),
		}
		_, _, err := producer.SendMessage(msg)
		if err != nil {
			fmt.Println("Error publish: ", err.Error())
		}

	}
}
