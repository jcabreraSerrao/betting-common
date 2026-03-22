package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

// ConsumerConfig contiene la configuración para crear un consumidor
type ConsumerConfig struct {
	Brokers []string
	GroupID string
	Topic   string   // Tópico simple (compatibilidad)
	Topics  []string // Lista de tópicos (nuevo)
}

// KafkaConsumer es un wrapper estructurado para segmentio/kafka-go
type KafkaConsumer struct {
	reader *kafka.Reader
}

// NewKafkaConsumer inicializa un nuevo consumidor
func NewKafkaConsumer(cfg ConsumerConfig) *KafkaConsumer {
	topics := cfg.Topics
	if len(topics) == 0 && cfg.Topic != "" {
		topics = []string{cfg.Topic}
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  cfg.Brokers,
		GroupID:  cfg.GroupID,
		GroupTopics: topics,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	return &KafkaConsumer{
		reader: r,
	}
}

// ReadMessage espera y lee el siguiente mensaje disponible en el tópico.
// Importante: No hace commit automático. Debes llamar a CommitMessage si fue exitoso.
func (c *KafkaConsumer) ReadMessage(ctx context.Context) (kafka.Message, error) {
	return c.reader.FetchMessage(ctx)
}

// CommitMessage marca el mensaje como procesado exitosamente (avanza el offset).
func (c *KafkaConsumer) CommitMessage(ctx context.Context, msg kafka.Message) error {
	return c.reader.CommitMessages(ctx, msg)
}

// Close finaliza la conexión del consumidor
func (c *KafkaConsumer) Close() error {
	if c.reader != nil {
		return c.reader.Close()
	}
	return nil
}
