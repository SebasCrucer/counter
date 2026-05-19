package coordinator

import (
	"coordinator/internal/indexer"
	"fmt"
	"net"
	"os"
	"sync"
	"sync/atomic"
)

type Coordinator struct {
	mu sync.Mutex
	Workers map[string]*Worker
	Indexer *indexer.Indexer
	Tasks chan int
	CompletedSteps int64
}

func NewCoordinator(file *os.File) *Coordinator {
	return &Coordinator{
		Workers: make(map[string]*Worker),
		Indexer: indexer.NewIndexer(file),
		Tasks: make(chan int, 10),
		CompletedSteps: 0,
	}
}

func (c *Coordinator) Init() *Coordinator {
	go func (){
		for i := range c.Indexer.Steps {
			c.Tasks<-i
		}
	}()

	return c
}

func (c *Coordinator) NewWorker(conn net.Conn) *Worker {
	c.mu.Lock()
	defer c.mu.Unlock()

	workerId := fmt.Sprintf("worker-%d", len(c.Workers)+1)

	worker := &Worker{
		Coordinator: c,
		ChunkIndex: <-c.Tasks,
		State: Working,
		Conn: conn,
	}

	c.Workers[workerId] = worker

	fmt.Printf("%s: task %d\n", workerId, worker.ChunkIndex)

	return worker
}

func (c *Coordinator) GetWorker(workerID string) (*Worker, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	worker, exists := c.Workers[workerID]
	return worker, exists
}

func (c *Coordinator) HadEnded() bool {
	return atomic.LoadInt64(&c.CompletedSteps) == int64(c.Indexer.Steps)
}