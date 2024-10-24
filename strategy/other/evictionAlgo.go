package other

type EvictionAlgo interface {
	evict(c *Cache)
}

type LRU struct{}

func (l LRU) evict(c *Cache) {
	// TODO
}

type LFU struct{}

func (l LFU) evict(c *Cache) {
	// TODO
}

type FIFO struct{}

func (f FIFO) evict(c *Cache) {
	// TODO
}
