package reporter

import (
	"errors"
	"fmt"
	"net"
	"worker/internal/protocol"
	"worker/internal/worker"
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

func (r *Reporter) Report() protocol.RCODE {
	conn, err := net.Dial("tcp", "localhost:3001")
	if err != nil {
		fmt.Printf("Error de conexión de reporter - worker %s: %s\n", r.Worker.WorkerId, err)
		return protocol.ROK
	}
	defer conn.Close()

	protocol.WriteRCOUNT(conn, r.Worker.Counter.Count)

	return protocol.ROK
}