package eventdriven

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type PubSub struct {
	conn        *amqp.Connection
	channel     *amqp.Channel
	subscribers []*Subscriber
}

func NewPubSub(amqpURL string) (*PubSub, error) {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return nil, err
	}

	return &PubSub{
		conn: conn,
	}, nil
}

type Event struct {
	ctx     context.Context
	Payload []byte
}

func (e *Event) Context() context.Context {
	if e.ctx == nil {
		return context.Background()
	}

	return e.ctx
}

func (s *PubSub) Subscribe(topic string, handler func(event *Event) error) error {
	channel, err := s.conn.Channel()
	if err != nil {
		return err
	}

	subscriber := &Subscriber{
		handler: handler,
		channel: channel,
	}

	queue, err := subscriber.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	subscriber.queue = queue
	err = subscriber.channel.QueueBind(queue.Name, "", topic, false, nil)
	if err != nil {
		return err
	}

	s.subscribers = append(s.subscribers, subscriber)

	return nil
}

func (s *PubSub) NewPublisher() (*Publisher, error) {
	channel, err := s.conn.Channel()
	if err != nil {
		return nil, err
	}

	return &Publisher{channel: channel}, nil
}

func (s *PubSub) Run(ctx context.Context) error {
	fmt.Println("PubSub is running...")
	for _, subscriber := range s.subscribers {
		// TODO: refactor this
		go func() {
			msgs, err := subscriber.channel.ConsumeWithContext(
				ctx,
				subscriber.queue.Name,
				"",
				true,
				false,
				false,
				false,
				nil,
			)
			if err != nil {
				fmt.Println("error", err)
				return
			}

			for msg := range msgs {
				err = subscriber.handler(&Event{Payload: msg.Body})
				if err != nil {
					nackErr := msg.Nack(false, true)
					if nackErr != nil {
						fmt.Println("Error nacking event", nackErr)
					}
					continue
				}

				// Auto-ack is already enabled
				// ackErr := msg.Ack(false)
				// if ackErr != nil {
				// 	fmt.Println("Error acking event", ackErr)
				// }
			}
		}()
	}

	return nil
}

func (s *PubSub) Close() error {
	for _, subscriber := range s.subscribers {
		err := subscriber.Close()
		if err != nil {
			return fmt.Errorf("error closing subscriber: %v", err)
		}
	}

	err := s.conn.Close()
	if err != nil {
		return fmt.Errorf("error closing connection: %v", err)
	}

	return nil
}
