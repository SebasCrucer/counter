package collector

import (
	"bufio"
	"coordinator/internal/collector"
	"coordinator/internal/protocol"
	"fmt"
	"net"
)

func HandleConnection(conn net.Conn, collector *collector.Collector) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	scanner.Split(bufio.ScanWords)
 	for scanner.Scan() {
		err := scanner.Err()
		if nil != err {
			fmt.Printf("Error de scanner del collector: %s\n", err)
			protocol.WriteRERROR(conn)
		}
		word := scanner.Text()

		var count int64
		n, err := fmt.Sscanf(word, string(protocol.RCOUNT), &count)
		if err == nil && n == 1 {
			collector.Collect(count)
			protocol.WriteROK(conn)
		}
 	}
	protocol.WriteRERROR(conn)
}
