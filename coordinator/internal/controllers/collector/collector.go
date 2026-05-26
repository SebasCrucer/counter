package collector

import (
	"coordinator/internal/collector"
	"coordinator/internal/coordinator"
	"coordinator/internal/utils"
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

	done := make(chan struct{})

	col := collector.NewCollector(c.Indexer.Steps, func(m map[string]uint64) {
		if err := utils.DumpToFile("assets/wordcount.tsv", m); err != nil {
			fmt.Println("Error escribiendo archivo:", err)
		} else {
			fmt.Printf("Resultado escrito: %d palabras\n", len(m))
		}
		close(done)
		ln.Close()
	})

	for {
		conn, err := ln.Accept()
		if err != nil {
			select {
			case <-done:
				fmt.Println("Todos los chunks fueron recibidos")
				return
			default:
				fmt.Println("Error al aceptar la conexión:", err)
				continue
			}	
		}
		fmt.Println("Nueva conexión de reporte")
		go HandleConnection(conn, col)
	}
}