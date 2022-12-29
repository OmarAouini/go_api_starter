package kafka

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

func ConsumeMessages(brokers []string, topic string, callBackForMessage func([]byte)) {
	worker, err := connectConsumer(brokers)
	if err != nil {
		panic(err)
	}
	logrus.Info("consuming partition soon...")

	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		logrus.Error(err)
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	// Count how many message processed
	msgCount := 0

	// Get signal for finish
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				logrus.Error(err)
			case msg := <-consumer.Messages():
				msgCount++
				logrus.Infof("Received message Count %d: | Topic(%s) | Message(%s) \n", msgCount, string(msg.Topic), string(msg.Value))
				callBackForMessage(msg.Value)
			case <-sigchan:
				logrus.Info("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	logrus.Info("Processed", msgCount, "messages")

	if err := worker.Close(); err != nil {
		logrus.Fatal(err)
	}
}

func connectConsumer(brokers []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	conn, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	logrus.Info("consumer ready")
	return conn, nil
}
