package other

type Cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo
	capacity     int
	maxCapacity  int
}

func InitCache(e EvictionAlgo) *Cache {
	store := make(map[string]string)
	return &Cache{
		storage:      store,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *Cache) SetEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

func (c *Cache) Add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evictionAlgo.evict(c)
		c.capacity--
	}
	c.capacity++
	c.storage[key] = value
}

func main() {
	evicLfu := LFU{}
	cache := InitCache(evicLfu)
	cache.Add("key1", "value1")

	evicFifo := FIFO{}
	cache.SetEvictionAlgo(evicFifo)
	cache.Add("key2", "value2")
}
