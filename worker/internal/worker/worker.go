package worker

import (
	"bufio"
	"fmt"
	"net"
	"sync/atomic"
	"worker/internal/counter"
	"worker/internal/protocol"
)

type Worker struct {
	WorkerId string
	TopJobs *int64
	Counter *counter.Counter
}

func NewWorker(w *Worker) protocol.WCODE {
	conn, err := net.Dial("tcp", "localhost:3000")
	id := w.WorkerId
	if err != nil {
		fmt.Printf("Error de conexión del worker %s: %s\n", id, err)
		return protocol.WERROR
	}
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	scanner.Split(bufio.ScanWords)
 	for scanner.Scan() {
		errS := scanner.Err()
		if nil != errS {
			fmt.Printf("Error de scanner del worker %s: %s\n", id, err)
			return protocol.WERROR
		}
		word := scanner.Text()

		if word == protocol.GENDCODE {
			return protocol.WGENDCODE
		}
		w.Counter.AddWord(&word)
 	}

	fmt.Printf("Worker: %s - Counted: %d - Top Jobs: %d \n", 
		id, 
		w.Counter.Count, 
		atomic.LoadInt64(w.TopJobs),
	)
	conn.Close()
	return protocol.WOK
}