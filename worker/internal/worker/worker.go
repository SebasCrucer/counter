package worker

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
	"sync/atomic"
	"unicode"
	"worker/internal/counter"
	"worker/internal/protocol"
)

type WCODE int

const (
	_ WCODE = iota
	WOK
	WGENDCODE
	WERROR
)

type Worker struct {
	WorkerId string
	TopJobs *int64
	Counter *counter.Counter
	ChunkId uint32
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

func (w *Worker) Work() WCODE {
	conn, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		fmt.Printf("Error de conexión del worker %s: %s\n", w.WorkerId, err)
		return WERROR
	}
	defer conn.Close()

	var greet [1]byte
	if _, err := io.ReadFull(conn, greet[:]); err != nil {
		return WERROR
	}

	if protocol.PROTOCOLCODE(greet[0]) == protocol.WGENDCODE {
		fmt.Printf("Worker %s: no hay trabajo\n", w.WorkerId)
		return WGENDCODE
	}

	var chunkBuf [4]byte
	if _, err := io.ReadFull(conn, chunkBuf[:]); err != nil {
		return WERROR
	}

	w.ChunkId = binary.BigEndian.Uint32(chunkBuf[:])

	scanner := bufio.NewScanner(conn)
	scanner.Buffer(make([]byte, 64*1024), 16*1024*1024)
	scanner.Split(bufio.ScanWords)
 	for scanner.Scan() {
		word := scanner.Text()

		tokens := strings.FieldsFunc(word, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsDigit(r)
		})

		for _, cleanWord := range tokens {
			w.Counter.AddWord(strings.ToLower(cleanWord))
		}
 	}

	errS := scanner.Err()
	if nil != errS {
		fmt.Printf("Error de scanner del worker %s: %s\n", w.WorkerId, errS)
		return WERROR
	}

	fmt.Printf("Worker: %s - Unique Words: %d - Top Jobs: %d \n", 
		w.WorkerId, 
		len(w.Counter.Count), 
		atomic.LoadInt64(w.TopJobs),
	)
	return WOK
}