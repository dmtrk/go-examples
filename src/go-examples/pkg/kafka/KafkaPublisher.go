package kafka

import (
	"github.com/Shopify/sarama"
	"go-examples/pkg/util"
	"fmt"
	"errors"
	"strings"
)

const (
	DEFAULT_BROKERS = "localhost:9092"
	DEFAULT_RETRY_MAX = 5
)

type KafkaPublisher struct {
	config   *sarama.Config
	producer sarama.SyncProducer
	//
	Brokers  []string
	Topic    string
}

func NewKafkaPublisher(properties map[string]string) *KafkaPublisher {
	instance := new(KafkaPublisher)
	instance.config = sarama.NewConfig()
	instance.config.Producer.Return.Successes = true //must be true to be used in a SyncProducer
	instance.config.Producer.RequiredAcks = sarama.WaitForAll
	instance.config.Producer.Retry.Max = util.GetInt(properties, "kafka.retry_max", DEFAULT_RETRY_MAX)
	//
	instance.Brokers = strings.Split(util.GetStr(properties, "kafka.brokers", DEFAULT_BROKERS), ",")
	instance.Topic = util.GetStr(properties, "kafka.topic", "")
	//
	return instance
}

func (self *KafkaPublisher) IsConnected() bool {
	return self.producer != nil
}

func (self *KafkaPublisher) Connect() error {
	shutdown(self.producer)
	p, err := sarama.NewSyncProducer(self.Brokers, self.config)
	if err != nil {
		return err
	}
	self.producer = p
	//
	return nil
}

func (self *KafkaPublisher) Disconnect() {
	fmt.Println("Disconnect()")
	shutdown(self.producer)
	self.producer = nil;
}

func (self *KafkaPublisher) Publish(headers map[string]string, data []byte) (error) {
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


		/*		err := self.channel.Publish(
					self.Exchange, // exchange
					self.RoutingKey, // routing key
					false, // mandatory
					false, // immediate
					amqp.Publishing{
						ContentType: "application/octet-stream",
						Body:        data,
					})*/
		return err
	}
}

func shutdown(producer sarama.SyncProducer) {
	if producer != nil {
		producer.Close()
	}
}