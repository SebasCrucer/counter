package coordinator

import (
	"coordinator/internal/coordinator"
	"coordinator/internal/protocol"
	"fmt"
	"net"
)

func HandleConnection(conn net.Conn, coordinator *coordinator.Coordinator) {
	defer conn.Close()
	fmt.Println("Worker conectado:", conn.RemoteAddr().String())
	worker := coordinator.NewWorker(conn)
	protocol.WriteWGOCODE(conn, worker.ChunkIndex)
	worker.Work()
}
