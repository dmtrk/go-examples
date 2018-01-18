package rabbit

import "github.com/streadway/amqp"

const (
	RABBIT_DEFAULT_URL = "amqp://localhost"
	RABBIT_DEFAULT_USERNAME = "guest"
	RABBIT_DEFAULT_PASSWORD = "guest"
)

func DialRabbit(url string) (*amqp.Connection, *amqp.Channel, error) {
	var err error
	var conn *amqp.Connection
	var ch *amqp.Channel
	conn, err = amqp.Dial(url)
	if err == nil {
		ch, err = conn.Channel()
	}
	return conn, ch, err
}