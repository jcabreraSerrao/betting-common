package nats

import (
	"fmt"
	"log"
	"time"

	natslib "github.com/nats-io/nats.go"
)

// StreamConfig contiene la configuración para crear un Stream en JetStream
type StreamConfig struct {
	Name     string
	Subjects []string
}

// SetupStream asegura que un Stream y sus consumidores existan en NATS
func SetupStream(conn *NATSConnection, cfg StreamConfig) error {
	// Verificar si el stream existe
	stream, err := conn.JS.StreamInfo(cfg.Name)
	if err != nil && err != natslib.ErrStreamNotFound {
		return fmt.Errorf("error verificando stream %s: %w", cfg.Name, err)
	}

	if stream == nil {
		// Crear el stream si no existe
		_, err = conn.JS.AddStream(&natslib.StreamConfig{
			Name:     cfg.Name,
			Subjects: cfg.Subjects,
			Storage:  natslib.FileStorage,
		})
		if err != nil {
			return fmt.Errorf("error creando stream %s: %w", cfg.Name, err)
		}
	} else {
		// Actualizar sujetos si es necesario
		_, err = conn.JS.UpdateStream(&natslib.StreamConfig{
			Name:     cfg.Name,
			Subjects: cfg.Subjects,
			Storage:  natslib.FileStorage,
		})
		if err != nil {
			return fmt.Errorf("error actualizando stream %s: %w", cfg.Name, err)
		}
	}

	return nil
}

// SetupPullConsumer asegura que un consumidor Pull exista para un stream.
// Si el consumidor existe con un FilterSubject diferente, lo recrea usando DeliverNewPolicy.
func SetupPullConsumer(conn *NATSConnection, streamName, consumerName, filterSubject string) error {
	info, err := conn.JS.ConsumerInfo(streamName, consumerName)
	if err == nil {
		// El consumidor ya existe. Si el filtro es el mismo, no hacer nada.
		if info.Config.FilterSubject == filterSubject {
			return nil
		}
		log.Printf("[NATS] Actualizando consumidor %s: el filtro cambió de '%s' a '%s'", consumerName, info.Config.FilterSubject, filterSubject)
		// El filtro cambió: borrarlo para recrearlo con el nuevo filtro
		_ = conn.JS.DeleteConsumer(streamName, consumerName)
	} else if err != natslib.ErrConsumerNotFound {
		return fmt.Errorf("error verificando consumidor %s: %w", consumerName, err)
	}

	log.Printf("[NATS] Creando consumidor Pull %s en stream %s con filtro '%s' (DeliverAllPolicy)", consumerName, streamName, filterSubject)
	// Crear consumidor nuevo con DeliverAllPolicy para procesar mensajes pendientes
	_, err = conn.JS.AddConsumer(streamName, &natslib.ConsumerConfig{
		Durable:       consumerName,
		DeliverPolicy: natslib.DeliverAllPolicy,
		AckPolicy:     natslib.AckExplicitPolicy,
		MaxDeliver:    10,
		AckWait:       60 * time.Second,
		FilterSubject: filterSubject,
	})
	if err != nil {
		return fmt.Errorf("error creando consumidor %s: %w", consumerName, err)
	}
	return nil
}

// DeleteStream elimina un stream de JetStream
func DeleteStream(conn *NATSConnection, name string) error {
	err := conn.JS.DeleteStream(name)
	if err != nil && err != natslib.ErrStreamNotFound {
		return err
	}
	return nil
}

// DeleteConsumer elimina un consumidor de un stream
func DeleteConsumer(conn *NATSConnection, streamName, consumerName string) error {
	err := conn.JS.DeleteConsumer(streamName, consumerName)
	if err != nil && err != natslib.ErrConsumerNotFound {
		return err
	}
	return nil
}
