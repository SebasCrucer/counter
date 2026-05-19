package collector

import (
	"fmt"
	"sync/atomic"
)

type Collector struct {
	Sum *int64
}

func NewCollector() *Collector {
	var sum int64 = 0
	return &Collector{
		Sum: &sum,
	}
}

func (c *Collector) Collect(count int64) {
	fmt.Printf("Collector recibió: %d - Total: %d\n", count, atomic.LoadInt64(c.Sum))
	atomic.AddInt64(c.Sum, count)
}