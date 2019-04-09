package model

type Config struct {
	APPName string `default:"app name"`
	Kafka   struct {
		Brokers []string
		Topic   string
		Group   string
	}
}
