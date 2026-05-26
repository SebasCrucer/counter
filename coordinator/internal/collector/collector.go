package collector

import (
	"sync"
)

type Collector struct {
	mu sync.Mutex
	Collection map[string]uint64
	collected map[uint32]struct{}
	toCollect uint32
	onComplete func(map[string]uint64)
	once sync.Once
}

func NewCollector(toCollect uint32, onComplete func(map[string]uint64)) *Collector {
	return &Collector{
		Collection: make(map[string]uint64),
		collected: make(map[uint32]struct{}),
		toCollect: toCollect,
		onComplete: onComplete,
	}
}

func (c *Collector) Collect(word string, count uint64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Collection[word] = c.Collection[word] + count
	// fmt.Printf("Collector recibió: %d %s - Total: %d\n", count, word, c.Collection[word])
}

func (c *Collector) MarkChunk(chunkId uint32) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.collected[chunkId]; ok {
		return false
	}

	c.collected[chunkId] = struct{}{}

	if uint32(len(c.collected)) == c.toCollect {
		c.once.Do(func() {go c.onComplete(c.Collection)})
	}
	return true
}