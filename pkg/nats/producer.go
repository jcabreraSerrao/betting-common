package nats

import (
	"context"
	"fmt"
	"strings"

	natslib "github.com/nats-io/nats.go"
)

// NATSProducer implementa la lógica de publicación de mensajes en JetStream
type NATSProducer struct {
	js natslib.JetStreamContext
}

// NewNATSProducer inicializa un nuevo productor de NATS
func NewNATSProducer(conn *NATSConnection) *NATSProducer {
	return &NATSProducer{
		js: conn.JS,
	}
}

// PublishMessage publica un mensaje en NATS JetStream.
// Si se provee una 'key', se concatena al subject para soportar FIFO por subject (ej: subject.key).
func (p *NATSProducer) PublishMessage(ctx context.Context, subject string, key string, payload []byte) error {
	fullSubject := subject
	if key != "" {
		// Normalizar key para evitar caracteres inválidos en NATS subjects
		cleanKey := strings.ReplaceAll(key, ":", ".")
		cleanKey = strings.ReplaceAll(cleanKey, "@", ".")
		fullSubject = fmt.Sprintf("%s.%s", subject, cleanKey)
	}

	_, err := p.js.Publish(fullSubject, payload, natslib.Context(ctx))
	if err != nil {
		return fmt.Errorf("error publicando en NATS: %w", err)
	}

	return nil
}

// PublishMessageWithOpts permite publicación con opciones específicas de NATS
func (p *NATSProducer) PublishMessageWithOpts(ctx context.Context, subject string, payload []byte, opts ...natslib.PubOpt) error {
	_, err := p.js.Publish(subject, payload, append(opts, natslib.Context(ctx))...)
	return err
}
