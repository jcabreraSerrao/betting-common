package nats

import (
	"fmt"
	"log"
	"time"

	natslib "github.com/nats-io/nats.go"
)

// NATSConnection gestiona la conexión a NATS y JetStream
type NATSConnection struct {
	Conn *natslib.Conn
	JS   natslib.JetStreamContext
}

// NewNATSConnection establece una nueva conexión a NATS con JetStream habilitado
func NewNATSConnection(url, user, pass string) (*NATSConnection, error) {
	opts := []natslib.Option{
		natslib.Name("BettingSystem"),
		natslib.MaxReconnects(-1),
		natslib.ReconnectWait(2 * time.Second),
		natslib.DisconnectErrHandler(func(nc *natslib.Conn, err error) {
			log.Printf("Desconectado de NATS: %v", err)
		}),
		natslib.ReconnectHandler(func(nc *natslib.Conn) {
			log.Printf("Reconectado a NATS: %s", nc.ConnectedUrl())
		}),
	}

	if user != "" && pass != "" {
		opts = append(opts, natslib.UserInfo(user, pass))
	}

	nc, err := natslib.Connect(url, opts...)
	if err != nil {
		return nil, fmt.Errorf("error conectando a NATS: %w", err)
	}

	js, err := nc.JetStream()
	if err != nil {
		nc.Close()
		return nil, fmt.Errorf("error habilitando JetStream: %w", err)
	}

	return &NATSConnection{
		Conn: nc,
		JS:   js,
	}, nil
}

// Close cierra la conexión a NATS
func (n *NATSConnection) Close() {
	if n.Conn != nil {
		n.Conn.Close()
	}
}
