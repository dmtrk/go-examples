package kafka

import (
	"github.com/Shopify/sarama"
	"go-examples/pkg/util"
	"fmt"
	"strings"
)

type KafkaConsumer struct {
	config   *sarama.Config
	consumer sarama.Consumer
	//
	Brokers  []string
	Topic    string
}

func NewKafkaConsumer(properties map[string]string) *KafkaConsumer {
	instance := new(KafkaConsumer)
	instance.config = sarama.NewConfig()
	instance.config.Producer.Return.Successes = true //must be true to be used in a SyncProducer
	instance.config.Producer.RequiredAcks = sarama.WaitForAll
	instance.config.Producer.Retry.Max = util.GetInt(properties, "kafka.retry_max", KAFKA_DEFAULT_RETRY_MAX)
	//
	instance.Brokers = strings.Split(util.GetStr(properties, "kafka.brokers", KAFKA_DEFAULT_BROKERS), ",")
	instance.Topic = util.GetStr(properties, "kafka.topic", "")
	//
	return instance
}

func (self *KafkaConsumer) IsConnected() bool {
	return self.consumer != nil
}

func (self *KafkaConsumer) Connect() error {
	shutdownConsumer(self.consumer)
	p, err := sarama.NewConsumer(self.Brokers, self.config)
	if err != nil {
		return err
	}
	self.consumer = p
	fmt.Println("Connected() consumer: ", self.consumer)
	//
	return nil
}

func (self *KafkaConsumer) Disconnect() {
	fmt.Println("Disconnect()")
	shutdownConsumer(self.consumer)
	self.consumer = nil;
}

func shutdownConsumer(consumer sarama.Consumer) {
	if consumer != nil {
		consumer.Close()
	}
}

