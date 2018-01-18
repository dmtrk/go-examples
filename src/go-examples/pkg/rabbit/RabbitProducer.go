package rabbit

import (
	"github.com/streadway/amqp"
	"go-examples/pkg/util"
	"fmt"
	"errors"
)

type RabbitProducer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	//
	Url        string
	Exchange   string
	RoutingKey string
	Username   string
	Password   string
}

func NewRabbitProducer(properties map[string]string) *RabbitProducer {
	instance := new(RabbitProducer)
	//
	instance.Url = util.GetStr(properties, "amqp.url", RABBIT_DEFAULT_URL)
	instance.Exchange = util.GetStr(properties, "amqp.exchange", "")
	instance.RoutingKey = util.GetStr(properties, "amqp.routing_key", "")
	instance.Username = util.GetStr(properties, "amqp.username", RABBIT_DEFAULT_USERNAME)
	instance.Password = util.GetStr(properties, "amqp.password", RABBIT_DEFAULT_PASSWORD)
	//
	return instance
}

func (self *RabbitProducer) IsConnected() bool {
	return self.connection != nil && self.channel != nil
}

func (self *RabbitProducer) Connect() error {
	shutdownProducer(self.connection, self.channel)
	conn, err := amqp.Dial(self.Url)
	if err != nil {
		return err
	}
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	self.connection = conn
	self.channel = ch
	//
	return nil
}

func (self *RabbitProducer) Disconnect() {
	fmt.Println("Disconnect()")
	shutdownProducer(self.connection, self.channel)
	self.connection = nil;
	self.channel = nil;
}

func (self *RabbitProducer) Publish(headers map[string]string, data []byte) (error) {
	fmt.Println("Publish()")
	if !self.IsConnected() {
		return errors.New("Not connected")
	} else {
		err := self.channel.Publish(
			self.Exchange, // exchange
			self.RoutingKey, // routing key
			false, // mandatory
			false, // immediate
			amqp.Publishing{
				ContentType: "application/octet-stream",
				Body:        data,
			})
		return err
	}
}

func shutdownProducer(connection *amqp.Connection, channel *amqp.Channel) {
	if channel != nil {
		channel.Close()
	}
	if connection != nil {
		connection.Close()
	}
}