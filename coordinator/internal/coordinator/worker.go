package coordinator

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"sync/atomic"
)

type State int

const (
	_ State = iota
	Working
	Finished
)

type Worker struct {
	Coordinator *Coordinator
	ChunkIndex uint32
	State State
	Conn net.Conn
}

func (w *Worker) Finish() {
	w.State = Finished
}

func (w *Worker) GetFileSection() *io.SectionReader {
	chunk := w.Coordinator.Indexer.IndexChunk(w.ChunkIndex)
	nextChunk := w.Coordinator.Indexer.IndexChunk(w.ChunkIndex+1)

	chunkSize := nextChunk.StartPosition - chunk.StartPosition

	fmt.Printf("ChunkIndex: %d Start: %d Size:%d\n", w.ChunkIndex, chunk.StartPosition, chunkSize)

	return io.NewSectionReader(
		w.Coordinator.Indexer.File, 
		chunk.StartPosition,
		chunkSize,
	)
}

func (w *Worker) Work() {
 	section := w.GetFileSection()
 	writter := bufio.NewWriter(w.Conn)

	buf := make([]byte, 32*1024)

	var failed bool
	for {
		n, errR := section.Read(buf)
		if n > 0 {
			_, errW := writter.Write(buf[:n])
			if errW != nil {
				fmt.Printf("Error al escribir Chunk %d: %s\n", w.ChunkIndex, errW)
				failed = true
				break
			}
		}

		if errR == io.EOF {
			break
		}
		if errR != nil {
			fmt.Printf("Error al leer Chunk %d: %s\n", w.ChunkIndex, errR)
			failed = true
			break
		}
 	}

	if errF := writter.Flush(); errF != nil {
		failed = true
	}

	if failed {
		w.OnError()
		return
	}
	w.OnSuccess()
} 

func (w *Worker) OnError() {
	w.Coordinator.Tasks<-w.ChunkIndex
}

func (w *Worker) OnSuccess() {
	atomic.AddInt64(&w.Coordinator.CompletedSteps, 1)
}