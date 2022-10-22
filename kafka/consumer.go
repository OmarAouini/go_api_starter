package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

/*
func ConsumeMessages(brokers []string, topic string) {
	worker, err := connectConsumer(brokers)
	if err != nil {
		panic(err)
	}
	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
}
*/

func connectConsumer(brokers []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	conn, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return conn, nil
}
