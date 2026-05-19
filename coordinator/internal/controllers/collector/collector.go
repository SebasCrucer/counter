package collector

import (
	"coordinator/internal/collector"
	"fmt"
	"net"
	"os"
)

func Collect() {
	ln, err := net.Listen("tcp", ":3001")
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
		os.Exit(1)
	}
	defer ln.Close()

	c := collector.NewCollector()

	for {
		conn, err := ln.Accept()
		fmt.Println("Nueva conexión entrante...")
		if err != nil {
			fmt.Println("Error al aceptar la conexión:", err)
			continue
		}
		go HandleConnection(conn, c)
	}
}