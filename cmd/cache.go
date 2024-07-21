package cmd

import "sync"

type Cache struct {
	store map[string]string
	mutex sync.Mutex
}

func NewCache() (cache *Cache) {
	cache = &Cache{}
	cache.store = make(map[string]string)
	return cache
}

func (cache *Cache) fetch(key string) (string, bool) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	value, ok := cache.store[key]
	return value, ok
}

func (cache *Cache) persist(key, value string) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	cache.store[key] = value
}
