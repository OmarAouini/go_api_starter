package main

import (
	"time"

	"github.com/OmarAouini/go_tdd/config"
	"github.com/OmarAouini/go_tdd/kafka"
	"github.com/OmarAouini/go_tdd/logging"
)

type Person struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

func main() {
	logging.ConfigureLogger()
	config.LoadEnv()

	brokers := []string{"localhost:29092"}
	kafka.SendMessage(brokers, "test_topic", Person{ID: 22, Nome: "test_name"})
	time.Sleep(time.Second * 3)
	kafka.ConsumeMessages(brokers, "test_topic", func(b []byte) {
		//action here
	})
}
