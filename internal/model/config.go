package model

type Config struct {
	APPName string `default:"app name"`
	Kafka   struct {
		Brokers       []string
		ConsumerTopic string
		ProducerTopic string
		Group         string
	}
}
