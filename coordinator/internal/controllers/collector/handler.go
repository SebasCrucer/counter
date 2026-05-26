package collector

import (
	"bufio"
	"coordinator/internal/collector"
	"coordinator/internal/protocol"
	"errors"
	"fmt"
	"io"
	"net"
)

func HandleConnection(conn net.Conn, collector *collector.Collector) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	var fatal error
 	for {
		word, count, err := protocol.ReadReport(reader)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Printf("Error leyendo reporte: %s\n", err)
			fatal = err
			break
		}

		collector.Collect(word, count)
	}

	if fatal != nil {
		protocol.WriteRERROR(conn)
	} else {
		protocol.WriteROK(conn)
	}

}
