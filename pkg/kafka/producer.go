package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

// ProducerConfig contiene la configuración para crear un productor
type ProducerConfig struct {
	Brokers []string
	Topic   string
}

// KafkaProducer es un wrapper estructurado para enviar mensajes
type KafkaProducer struct {
	writer *kafka.Writer
}

// NewKafkaProducer inicializa un nuevo productor para un tópico específico
func NewKafkaProducer(cfg ProducerConfig) *KafkaProducer {
	w := &kafka.Writer{
		Addr:                   kafka.TCP(cfg.Brokers...),
		Topic:                  cfg.Topic,
		Balancer:               &kafka.Hash{}, // Vital para usar Partition Keys equitativamente
		AllowAutoTopicCreation: true,
	}

	return &KafkaProducer{
		writer: w,
	}
}

// PublishMessage publica un mensaje en Kafka asegurando que se asigne a la
// partición correcta basada en la 'key' (vital para FIFO estricto).
func (p *KafkaProducer) PublishMessage(ctx context.Context, key string, payload []byte) error {
	msg := kafka.Message{
		Value: payload,
	}
	// Aplicar la Partition Key si se provee
	if key != "" {
		msg.Key = []byte(key)
	}

	return p.writer.WriteMessages(ctx, msg)
}

// Close finaliza la conexión del productor
func (p *KafkaProducer) Close() error {
	if p.writer != nil {
		return p.writer.Close()
	}
	return nil
}
