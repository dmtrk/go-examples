package rabbit

import (
	"github.com/streadway/amqp"
	"go-examples/pkg/util"
	"fmt"
	"errors"
	"log"
)

type MessageListener interface {
	GetId() string
	OnMessage(headers map[string]string, data []byte) error
}

type RabbitClient struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	listener   MessageListener
	//
	Url        string
	Exchange   string
	RoutingKey string
	Queue      string
	Username   string
	Password   string
}

func NewRabbitClient(properties map[string]string) *RabbitClient {
	instance := new(RabbitClient)
	//
	instance.Url = util.GetStr(properties, "amqp.url", RABBIT_DEFAULT_URL)
	instance.Exchange = util.GetStr(properties, "amqp.exchange", "")
	instance.RoutingKey = util.GetStr(properties, "amqp.routing_key", "")
	instance.Queue = util.GetStr(properties, "amqp.queue", "")
	instance.Username = util.GetStr(properties, "amqp.username", RABBIT_DEFAULT_USERNAME)
	instance.Password = util.GetStr(properties, "amqp.password", RABBIT_DEFAULT_PASSWORD)
	//
	return instance
}

func (self *RabbitClient) IsConnected() bool {
	return self.connection != nil && self.channel != nil
}

func (self *RabbitClient) Connect() error {
	shutdown(self.connection, self.channel)
	//
	conn, ch, err := DialRabbit(self.Url)
	if err == nil {
		self.connection = conn
		self.channel = ch
	}
	return err
}

func (self *RabbitClient) Disconnect() {
	fmt.Println("Disconnect()")
	shutdown(self.connection, self.channel)
	self.connection = nil;
	self.channel = nil;
}

func (self *RabbitClient) Publish(headers map[string]string, data []byte) (error) {
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

func (self *RabbitClient) Consume(listener MessageListener) error {
	fmt.Println("Consume()")
	if !self.IsConnected() {
		return errors.New("Not connected")
	}else if listener==nil {
		return errors.New("MessageListener must not be nil")
	}else{
		if(self.listener!=nil){
			fmt.Println("Consume() replacing listener:",self.listener)
		}
		self.listener = listener
		msgs, err := self.channel.Consume(
			self.Queue, // queue
			self.listener.GetId(), // consumer-id
			false,   // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // args
		)
		if err==nil {
			go func() {
				for msg := range msgs {
					log.Printf("Received a message: %s", msg.DeliveryTag)
					// copy headers to map
					headers := make(map[string]string, len(msg.Headers))
					for key,val := range msg.Headers{
						headers[key]=fmt.Sprintf("%v",val)
					}
					// call OnMessage()
					if self.listener.OnMessage(headers,msg.Body)==nil {
						log.Printf("Ack message: %s", msg.DeliveryTag)
						msg.Acknowledger.Ack(msg.DeliveryTag, false)
					}else{
						log.Printf("Reject message: %s", msg.DeliveryTag)
						msg.Acknowledger.Reject(msg.DeliveryTag, false)
					}
				}
			}()
		}
		return err
	}
}

func shutdown(connection *amqp.Connection, channel *amqp.Channel) {
	if channel != nil {
		channel.Close()
	}
	if connection != nil {
		connection.Close()
	}
}
