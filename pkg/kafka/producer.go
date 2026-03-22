package kafka

import (
	"context"
	"fmt"

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
		Balancer:               &kafka.Hash{}, // Vital para usar Partition Keys equitativamente
		AllowAutoTopicCreation: true,
	}

	return &KafkaProducer{
		writer: w,
	}
}

// PublishMessage publica un mensaje en Kafka asegurando que se asigne a la
// partición correcta basada en la 'key' (vital para FIFO estricto).
func (p *KafkaProducer) PublishMessage(ctx context.Context, topic string, key string, payload []byte) error {
	fmt.Printf("[KAFKA-PRODUCER] Publicando en tópico: %s, key: %s, size: %d\n", topic, key, len(payload))
	msg := kafka.Message{
		Topic: topic,
		Value: payload,
	}
	// Aplicar la Partition Key si se provee
	if key != "" {
		msg.Key = []byte(key)
	}

	err := p.writer.WriteMessages(ctx, msg)
	if err != nil {
		fmt.Printf("[KAFKA-PRODUCER] ERROR: %v\n", err)
	} else {
		fmt.Printf("[KAFKA-PRODUCER] Publicado OK\n")
	}
	return err
}

// PublishMessageToTopic es ahora un alias de PublishMessage para mantener compatibilidad si se usa en otros sitios.
func (p *KafkaProducer) PublishMessageToTopic(ctx context.Context, topic string, key string, payload []byte) error {
	return p.PublishMessage(ctx, topic, key, payload)
}

// Close finaliza la conexión del productor
func (p *KafkaProducer) Close() error {
	if p.writer != nil {
		return p.writer.Close()
	}
	return nil
}
