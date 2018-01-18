package rabbit

import (
	"github.com/streadway/amqp"
	"go-examples/pkg/util"
	"fmt"
)

type RabbitConsumer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	//
	Url        string
	Queue   string
	Username   string
	Password   string
}

func NewRabbitConsumer(properties map[string]string) *RabbitConsumer {
	instance := new(RabbitConsumer)
	//
	instance.Url = util.GetStr(properties, "amqp.url", RABBIT_DEFAULT_URL)
	instance.Queue = util.GetStr(properties, "amqp.queue", "")
	instance.Username = util.GetStr(properties, "amqp.username", RABBIT_DEFAULT_USERNAME)
	instance.Password = util.GetStr(properties, "amqp.password", RABBIT_DEFAULT_PASSWORD)
	//
	return instance
}

func (self *RabbitConsumer) IsConnected() bool {
	return self.connection != nil && self.channel != nil
}

func (self *RabbitConsumer) Connect() error {
	shutdownConsumer(self.connection, self.channel)
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


func (self *RabbitConsumer) Disconnect() {
	fmt.Println("Disconnect()")
	shutdownConsumer(self.connection, self.channel)
	self.connection = nil;
	self.channel = nil;
}


func shutdownConsumer(connection *amqp.Connection, channel *amqp.Channel) {
	if channel != nil {
		channel.Close()
	}
	if connection != nil {
		connection.Close()
	}
}

