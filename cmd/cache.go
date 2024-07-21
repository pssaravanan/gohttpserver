package cmd

type Cache struct {
	store map[string]string
}

func NewCache() (cache *Cache) {
	cache = &Cache{}
	cache.store = make(map[string]string)
	return cache
}

func (cache *Cache) fetch(key string) (string, bool) {
	value, ok := cache.store[key]
	return value, ok
}

func (cache *Cache) persist(key, value string) {
	cache.store[key] = value
}
