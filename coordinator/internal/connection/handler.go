package connection

import (
	"coordinator/internal/coordinator"
	"fmt"
	"net"
)

func HandleConnection(conn net.Conn, coordinator *coordinator.Coordinator) {
	defer conn.Close()
	fmt.Println("Worker conectado:", conn.RemoteAddr().String())
	worker := coordinator.NewWorker(conn)
	worker.Work()
}
