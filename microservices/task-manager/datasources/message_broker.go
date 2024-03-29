package datasources

import (
	"fabc.it/task-manager/config"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

const (
	NewSubtasksQueue       = "newSubtasks"
	CompletedSubtasksQueue = "completedSubtasks"
)

type MessageBroker struct {
	*amqp.Channel
}

func NewMessageBroker(env *config.Env) *MessageBroker {
	uri := fmt.Sprintf("amqp://%s:%s@%s:%s", env.RabbitMQUsername, env.RabbitMQPassword, env.RabbitMQHost, env.RabbitMQPort)

	conn, err := amqp.Dial(uri)
	if err != nil {
		log.Fatal(err)
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	// Declare queues
	_, err = channel.QueueDeclare(
		NewSubtasksQueue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	_, err = channel.QueueDeclare(
		CompletedSubtasksQueue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	return &MessageBroker{
		channel,
	}
}
