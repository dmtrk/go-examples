package kafka

import (
	"github.com/Shopify/sarama"
	"go-examples/pkg/util"
	"fmt"
	"errors"
	"strings"
)

type KafkaProducer struct {
	config   *sarama.Config
	producer sarama.SyncProducer
	//
	Brokers  []string
	Topic    string
}

func NewKafkaProducer(properties map[string]string) *KafkaProducer {
	instance := new(KafkaProducer)
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

func (self *KafkaProducer) IsConnected() bool {
	return self.producer != nil
}

func (self *KafkaProducer) Connect() error {
	shutdownProducer(self.producer)
	p, err := sarama.NewSyncProducer(self.Brokers, self.config)
	if err != nil {
		return err
	}
	self.producer = p
	fmt.Println("Connected() producer: ", self.producer)
	//
	return nil
}

func (self *KafkaProducer) Disconnect() {
	fmt.Println("Disconnect()")
	shutdownProducer(self.producer)
	self.producer = nil;
}

func (self *KafkaProducer) Publish(headers map[string]string, data []byte) (error) {
	fmt.Println("Publish()")
	if !self.IsConnected() {
		return errors.New("Not connected")
	} else {
		msg := &sarama.ProducerMessage{
			Topic: self.Topic,
			Value: sarama.ByteEncoder(data),
		}
		partition, offset, err := self.producer.SendMessage(msg)
		fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", self.Topic, partition, offset)
		return err
	}
}

func shutdownProducer(producer sarama.SyncProducer) {
	if producer != nil {
		producer.Close()
	}
}