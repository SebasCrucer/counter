package coordinator

import (
	"coordinator/internal/coordinator"
	"coordinator/internal/protocol"
	"fmt"
	"net"
	"os"
)

func Coordinate(file *os.File) {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
		os.Exit(1)
	}
	defer ln.Close()

	c := coordinator.NewCoordinator(file).Init()

	for {
		conn, err := ln.Accept()
		if c.HadEnded() {
			protocol.WriteGENDCODE(conn)
			conn.Close()
			return
		}
		fmt.Println("Nueva conexión entrante...")
		if err != nil {
			fmt.Println("Error al aceptar la conexión:", err)
			continue
		}
		go HandleConnection(conn, c)
	}
}