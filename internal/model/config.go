package model

import "fmt"

type Config struct {
	APPName string `default:"app name"`
	Kafka   struct {
		Brokers  []string
		Consumer []string
		Producer string
		Group    string
	}
}

func (config *Config) print() {
	fmt.Printf("%+v\n", config)
}
