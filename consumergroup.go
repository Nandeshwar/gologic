package main

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

type handler struct {
	chanSize int
}

func (h *handler) Setup(sarama.ConsumerGroupSession) error {
	logrus.Info("A Kafka consumer group session is starting")
	return nil
}

func (h *handler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	routineCnt := make(chan int, h.chanSize)
	warnLevel := int(math.Round(float64(h.chanSize) * .90))

	go func() {
		for {
			<-time.After(30 * time.Second)
			logChanLength(len(routineCnt), warnLevel)
		}
	}()

	/*
		for {
			select {
			case <-time.After(5 * time.Second):
				fmt.Println("Nothing came")
			case msg := <-claim.Messages():
				//fmt.Println("Received at Kafka messages", string(msg.Key), string(msg.Value))
				fmt.Println(string(msg.Value))
				sess.MarkMessage(msg, "")
			}
		}
	*/

	for message := range claim.Messages() {
		routineCnt <- 1
		go func(message *sarama.ConsumerMessage) {

			fmt.Println(string(message.Value))
			sess.MarkMessage(message, "")
			<-routineCnt
		}(message)

	}

	return nil

}

func (h *handler) Cleanup(sarama.ConsumerGroupSession) error {
	logrus.Info("A Kafka consumer group session is ending")
	return nil
}

func main() {
	brokers := []string{"localhost:9092"}
	consumerGroup, err := sarama.NewConsumerGroup(
		brokers,
		"radiot",
		NewConfig(),
	)
	if err != nil {
		logrus.WithError(err).Fatal("Error creating Kafka consumer")
	}

	h := &handler{chanSize: 250}

	go func() {
		for err := range consumerGroup.Errors() {
			logrus.WithError(err).Error("Error received from consumer group")
		}
	}()

	err = consumerGroup.Consume(context.Background(), []string{"dish-asset-telemetry"}, h)
	if err != nil {
		fmt.Println(err)
	}

	defer consumerGroup.Close()
}

// NewConfig returns a sarama config with our preferred defaults.
func NewConfig() *sarama.Config {
	c := sarama.NewConfig()
	c.Consumer.Return.Errors = true
	c.Version = sarama.V2_0_0_0
	return c
}

// logs warning if count is greater than 90% of max
// logs info if count is less than 90%
func logChanLength(length, warnLevel int) {
	if length >= warnLevel {
		logrus.WithFields(logrus.Fields{
			"count": length,
		}).Warn("Ingester goroutine count")
	}
}
