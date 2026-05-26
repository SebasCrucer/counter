package collector

import (
	"coordinator/internal/collector"
	"coordinator/internal/coordinator"
	"fmt"
	"net"
	"os"
)

func Collect(c *coordinator.Coordinator) {
	ln, err := net.Listen("tcp", ":3001")
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
		os.Exit(1)
	}
	defer ln.Close()

	col := collector.NewCollector()

	for {
		conn, err := ln.Accept()
		fmt.Println("Nueva conexión entrante...")
		if err != nil {
			fmt.Println("Error al aceptar la conexión:", err)
			continue
		}
		go HandleConnection(conn, col)
	}
}