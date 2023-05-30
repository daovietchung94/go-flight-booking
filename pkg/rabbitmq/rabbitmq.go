package rabbitmq

import (
	"fmt"
	"go-training/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewRabbitMQConn(config *config.Config) (*amqp.Connection, error) {
	connAddr := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		config.RabbitMQ.User,
		config.RabbitMQ.Password,
		config.RabbitMQ.Host,
		config.RabbitMQ.Port,
	)

	return amqp.Dial(connAddr)
}
