package nats

import (
	natslib "github.com/nats-io/nats.go"
)

// ErrTimeout se exporta para que los consumidores detecten tiempos de espera agotados
var ErrTimeout = natslib.ErrTimeout

// Message representa un mensaje genérico para NATS
type Message struct {
	Subject string
	Key     string
	Data    []byte
	Msg     *natslib.Msg // Referencia al mensaje original de NATS (para ACK)
}
