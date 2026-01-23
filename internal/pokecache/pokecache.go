package pokecache

import (
	"fmt"
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

func (cache *Cache) reapLoop() {
	ticker := time.NewTicker(5 * time.Second)

	// Create a loop to run at regular intervals
	for range ticker.C {
		// Loop through cache map
		for key, entry  := range cache.cacheEntry {
			// If key is createdAt, check the interval
			if key == "createdAt" {
				if time.Since(entry.createdAt) > cache.interval {
					fmt.Println("Clear this!")
				}
			}
		}
	}


}

// Create a new cache, pass an interval argument to decide how long objects in the cache reside
func NewCache(interval time.Duration) *Cache {
	
	cache := Cache{
		interval: interval,
		cacheEntry: make(map[string]CacheEntry),
	}
	// call reapLoop() method
	go cache.reapLoop()
	return &cache
}

func (cache *Cache) Add(key string, val []byte) {
	// create new entry
	entry := CacheEntry{
		createdAt: time.Now(),
		val: val,
	}

	// add to cache
	cache.cacheEntry[key] = entry
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	entry, ok := cache.cacheEntry[key]

	if ok {
		fmt.Println("Found value : " + string(entry.val))
		return entry.val, ok
	}

	return nil, ok
}
