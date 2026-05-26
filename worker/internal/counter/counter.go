package counter


type Counter struct {
	Count map[string]uint64
}

func NewCounter() *Counter {
	return &Counter{
		Count: make(map[string]uint64),
	}
}

func (c *Counter) AddWord(word string) {
	c.Count[word]++
}