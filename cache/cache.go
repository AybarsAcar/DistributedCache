package cache

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// implementation file

type Cache struct {
	lock sync.RWMutex
	data map[string][]byte
}

func New() *Cache {
	return &Cache{
		data: make(map[string][]byte),
	}
}

func (c *Cache) Get(key []byte) ([]byte, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	keyStr := string(key)

	val, ok := c.data[keyStr]

	if !ok {
		return nil, fmt.Errorf("key (%s) not found", keyStr)
	}

	log.Printf("GET %s = %s", string(key), string(val))

	return val, nil
}

func (c *Cache) Has(key []byte) bool {
	c.lock.RLock()
	defer c.lock.RUnlock()

	keyStr := string(key)

	_, ok := c.data[keyStr]

	return ok
}

// Set up the data then create a ticker on another thread then delete it
func (c *Cache) Set(key, value []byte, ttl time.Duration) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.data[string(key)] = value

	log.Printf("SET %s to %s", string(key), string(value))

	go func() {
		<-time.After(ttl)
		delete(c.data, string(key))
	}()

	return nil
}

func (c *Cache) Delete(key []byte) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	delete(c.data, string(key))

	return nil
}
