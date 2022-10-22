package kafka

import (
	"encoding/json"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

// Send message to kafka brokers on the specified topic, the message must be Json serializable
func SendMessage(brokers []string, topic string, message interface{}) error {

	messageBytes, err := json.Marshal(message)
	if err != nil {
		logrus.Error(err)
		return err
	}

	producer, err := connectProducer(brokers)
	if err != nil {
		logrus.Error(err)
		return err
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(messageBytes),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		logrus.Error(err)
		return err
	}

	logrus.Info("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)

	return nil
}

func connectProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return producer, nil
}
