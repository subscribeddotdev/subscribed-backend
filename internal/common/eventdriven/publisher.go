package eventdriven

import (
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	channel *amqp.Channel
}

func (p *Publisher) Publish(topic string, event *Event) error {
	err := p.channel.ExchangeDeclare(topic, "fanout", true, false, false, false, nil)
	if err != nil {
		return err
	}

	return p.channel.PublishWithContext(event.Context(), topic, "", false, false, amqp.Publishing{
		Headers:         nil,
		ContentType:     "application/json",
		ContentEncoding: "",
		DeliveryMode:    0,
		Priority:        0,
		CorrelationId:   "",
		ReplyTo:         "",
		Expiration:      "",
		MessageId:       "",
		Timestamp:       time.Now(),
		Type:            "",
		UserId:          "",
		AppId:           "",
		Body:            event.Payload,
	})
}

func (p *Publisher) Close() error {
	return p.channel.Close()
}
