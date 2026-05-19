package main

import (
	"coordinator/internal/connection"
	"coordinator/internal/coordinator"
	"coordinator/internal/protocol"
	"fmt"
	"net"
	"os"
)

func main() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
		os.Exit(1)
	}
	defer ln.Close()

	fmt.Println("Servidor TCP escuchando en el puerto 3000...")

	file, err := os.Open("assets/wiki_concatenated.txt")
 	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
 	}
 	defer file.Close()

	coordinator := coordinator.NewCoordinator(file)

	coordinator.Init()

	for {
		conn, err := ln.Accept()
		if coordinator.HadEnded() {
			protocol.WriteGENDCODE(conn)
			conn.Close()
			return
		}
		fmt.Println("Nueva conexión entrante...")
		if err != nil {
			fmt.Println("Error al aceptar la conexión:", err)
			continue
		}
		go connection.HandleConnection(conn, coordinator)
	}
}