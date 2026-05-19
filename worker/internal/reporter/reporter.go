package reporter

import (
	"fmt"
	"net"
	"worker/internal/protocol"
	"worker/internal/worker"
)


type Reporter struct {
	Worker *worker.Worker
}

func Report(r *Reporter) protocol.RCODE {
	conn, err := net.Dial("tcp", "localhost:3001")
	if err != nil {
		fmt.Printf("Error de conexión de reporter - worker %s: %s\n", r.Worker.WorkerId, err)
		return protocol.ROK
	}
	defer conn.Close()

	protocol.WriteRCOUNT(conn, r.Worker.Counter.Count)

	return protocol.ROK
}