package worker

import (
	"bufio"
	"errors"
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

func NewWorker(w *Worker) (*Worker, error){
	if w.TopJobs == nil {
		return nil, errors.New("topjobs is required")
	}
	if w.Counter == nil {
		return nil, errors.New("counter is required")
	}
	return w, nil
}

func (w *Worker) Work() protocol.WCODE {
	conn, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		fmt.Printf("Error de conexión del worker %s: %s\n", w.WorkerId, err)
		return protocol.WERROR
	}
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	scanner.Split(bufio.ScanWords)
 	for scanner.Scan() {
		errS := scanner.Err()
		if nil != errS {
			fmt.Printf("Error de scanner del worker %s: %s\n", w.WorkerId, err)
			return protocol.WERROR
		}
		word := scanner.Text()

		if word == string(protocol.GENDCODE) {
			return protocol.WGENDCODE
		}
		w.Counter.AddWord(&word)
 	}

	fmt.Printf("Worker: %s - Counted: %d - Top Jobs: %d \n", 
		w.WorkerId, 
		w.Counter.Count, 
		atomic.LoadInt64(w.TopJobs),
	)
	conn.Close()
	return protocol.WOK
}