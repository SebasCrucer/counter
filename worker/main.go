package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"worker/internal/protocol"
	"worker/internal/worker"
)

type Context struct {
	TOPJOB *int64
}

func main() {
	var TOPJOB int64 = 0

	jobs := make(chan int, 40)
	done := make(chan struct{})
	var onceDone sync.Once
	var onceJobs sync.Once
	var wg sync.WaitGroup
	for range 10 {
		wg.Go(func () {
			for jobId := range jobs {
				select {
				case<-done: return
				default:
				}
				id := fmt.Sprintf("J%d", jobId)

				w := worker.NewWorker(&worker.Worker{
					WorkerId: id,
					TopJobs: &TOPJOB,
				})
				if w == protocol.WGENDCODE {
					onceDone.Do(func(){close(done)})
					return
				}
			}
		})
	}
	go func() {
		for {
			select {
			case jobs<-int(atomic.LoadInt64(&TOPJOB)): atomic.AddInt64(&TOPJOB, 1)
			case <-done: onceJobs.Do(func(){close(jobs)}); return
			}
			
		}
	}()
	wg.Wait()
}