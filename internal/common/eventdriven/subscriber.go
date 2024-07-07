package eventdriven

import amqp "github.com/rabbitmq/amqp091-go"

type Subscriber struct {
	handler func(event *Event) error
	queue   amqp.Queue
	channel *amqp.Channel
	conn    *amqp.Connection
}

func (s *Subscriber) Close() error {
	return s.channel.Close()
}
