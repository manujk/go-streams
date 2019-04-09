package main

import (
	"fmt"
	"github.com/jinzhu/configor"
)

var Config = struct {
	APPName string `default:"app name"`
	Kafka   struct {
		Brokers []string
	}
}{}

func main() {
	_ = configor.Load(&Config, "./config/config.yml")
	fmt.Printf("config: %#v", Config.Kafka.Brokers)
}
