package vrmlgo

import "sync"

type Cache interface {
	Get(key string) (value string, found bool, err error)
	Set(key string, value string) error
}

// A simple in-memory cache store.
type LocalCache struct {
	sync.RWMutex

	store map[string]string
}

func NewLocalCache() *LocalCache {

	return &LocalCache{
		store: make(map[string]string),
	}
}

func (c *LocalCache) Get(key string) (string, bool, error) {
	c.RLock()
	defer c.RUnlock()

	v, ok := c.store[key]
	return v, ok, nil
}

func (c *LocalCache) Set(key string, value string) error {
	c.Lock()
	defer c.Unlock()

	c.store[key] = value

	return nil
}
