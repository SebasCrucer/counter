package coordinator

import (
	"coordinator/internal/coordinator"
	"coordinator/internal/protocol"
	"fmt"
	"net"
	"os"
)

func Coordinate(c *coordinator.Coordinator) {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
		os.Exit(1)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if c.HadEnded() {
			protocol.WriteWGENDCODE(conn)
			conn.Close()
			return
		}
		fmt.Println("Nueva conexión de worker")
		if err != nil {
			fmt.Println("Error al aceptar la conexión:", err)
			continue
		}
		go HandleConnection(conn, c)
	}
}