package main

import (
	"github.com/astomiwanda/go-coinbit/cmd/consumer"
	"github.com/astomiwanda/go-coinbit/cmd/publisher"
)

func main() {
	go consumer.RunConsumer()
	publisher.RubPublisher()
}
