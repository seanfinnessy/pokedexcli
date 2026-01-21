package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu         sync.Mutex            // protect cache across goroutines
	cacheEntry map[string]CacheEntry // entry for cache, containing created time and value
	interval   time.Duration
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

// Create a new cache, pass an interval argument to decde how long objects in the cache reside
func NewCache(interval time.Duration) *Cache {
	var cache Cache
	cache.mu = sync.Mutex{}
	cache.interval = interval

	return &cache
}

func (cache *Cache) Add(key string, val []byte) {

}

func (cache *Cache) Get(key string) ([]byte, bool) {

}
