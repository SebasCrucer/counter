package collector

import (
	"sync"
)

type Collector struct {
	mu sync.Mutex
	Collection map[string]uint64
}

func NewCollector() *Collector {
	return &Collector{
		Collection: make(map[string]uint64),
	}
}

func (c *Collector) Collect(word string, count uint64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Collection[word] = c.Collection[word] + count
	// fmt.Printf("Collector recibió: %d %s - Total: %d\n", count, word, c.Collection[word])
}