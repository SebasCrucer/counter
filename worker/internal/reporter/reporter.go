package reporter

import (
	"errors"
	"fmt"
	"io"
	"net"
	"worker/internal/protocol"
	"worker/internal/worker"
)

type RCODE int

const (
	_ RCODE = iota
	ROK
	RERROR
)

type Reporter struct {
	Worker *worker.Worker
}

func NewReporter(r *Reporter) (*Reporter, error) {
	if r.Worker == nil {
		return nil, errors.New("worker is required")
	}	
	return r, nil
}

func (r *Reporter) Report() RCODE {
	conn, err := net.Dial("tcp", "localhost:3001")
	if err != nil {
		fmt.Printf("Error de conexión de reporter - worker %s: %s\n", r.Worker.WorkerId, err)
		return RERROR
	}
	defer conn.Close()

	protocol.WriteRCOUNT(conn, r.Worker.Counter.Count)

	if tcp, ok := conn.(*net.TCPConn); ok {
		if err := tcp.CloseWrite(); err != nil {
			return RERROR
		}
	}

	var ack [1]byte
	if _, err := io.ReadFull(conn, ack[:]); err != nil {
		return RERROR
	}

	switch protocol.PROTOCOLCODE(ack[0]) {
	case protocol.ROK:
		return ROK
	case protocol.RERROR:
		return RERROR
	default: 
		return RERROR
	}
}