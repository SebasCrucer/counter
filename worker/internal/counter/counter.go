package counter


type Counter struct {
	Count int
}

func NewCounter() *Counter {
	return &Counter{
		Count: 0,
	}
}

func (c *Counter) AddWord(word *string) {
	c.Count++
}