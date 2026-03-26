package nats

import (
	"context"
	"fmt"
	"time"

	natslib "github.com/nats-io/nats.go"
)

// NATSConsumer implementa un consumidor de tipo Pull para JetStream
type NATSConsumer struct {
	sub *natslib.Subscription
}

// NewNATSConsumer inicializa un consumidor Pull en un stream específico
// consumerName debe estar previamente creado en NATS.
func NewNATSConsumer(conn *NATSConnection, subject, consumerName string) (*NATSConsumer, error) {
	sub, err := conn.JS.PullSubscribe(subject, consumerName)
	if err != nil {
		return nil, fmt.Errorf("error suscribiendo consumidor Pull: %w", err)
	}

	return &NATSConsumer{
		sub: sub,
	}, nil
}

// FetchMessage obtiene el siguiente mensaje disponible.
// Bloquea hasta que llegue un mensaje o se alcance el timeout.
func (c *NATSConsumer) FetchMessage(ctx context.Context, timeout time.Duration) (*Message, error) {
	if timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}

	msgs, err := c.sub.Fetch(1, natslib.Context(ctx))
	if err != nil {
		if err == context.DeadlineExceeded || err == natslib.ErrTimeout {
			return nil, ErrTimeout
		}
		return nil, err
	}

	if len(msgs) == 0 {
		return nil, natslib.ErrTimeout
	}

	m := msgs[0]
	return &Message{
		Subject: m.Subject,
		Data:    m.Data,
		Msg:     m,
	}, nil
}

// AckMessage marca el mensaje como procesado exitosamente
func (c *NATSConsumer) AckMessage(msg *Message) error {
	if msg.Msg != nil {
		return msg.Msg.Ack()
	}
	return nil
}

// NakMessage marca el mensaje como fallido para reintento (Negative Ack)
func (c *NATSConsumer) NakMessage(msg *Message) error {
	if msg.Msg != nil {
		return msg.Msg.Nak()
	}
	return nil
}

// NakWithDelay marca el mensaje como fallido para reintento con un retraso específico
func (c *NATSConsumer) NakWithDelay(msg *Message, delay time.Duration) error {
	if msg.Msg != nil {
		return msg.Msg.NakWithDelay(delay)
	}
	return nil
}

// Close cierra la suscripción del consumidor
func (c *NATSConsumer) Close() error {
	if c.sub != nil {
		return c.sub.Unsubscribe()
	}
	return nil
}
