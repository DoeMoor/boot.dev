package pokecache

import (
	"sync"
	"time"
)

type PokeCache struct {
	cache  map[string]cacheEntry
	mu     sync.RWMutex
	maxAge time.Duration
}

// createTime time.Time
// value      []byte // response body
type cacheEntry struct {
	createTime time.Time
	value      []byte // response body
}

var cache *PokeCache
var lock sync.Mutex

// GetCache returns a pointer to the cache instance.
// If the cache instance does not exist, it will create a new one.
// The cache instance is a singleton.
// The cache instance runs a goroutine in the background to remove expired entries.
func GetCache() *PokeCache {
	if cache == nil {
		lock.Lock()
		defer lock.Unlock()
		if cache == nil {
			cache = &PokeCache{
				cache:  make(map[string]cacheEntry),
				maxAge: 5 * time.Second,
			}
			go cache.cacheMaintenance()
		}
	}
	return cache
}

// Write stores the value associated with the key in the cache.
// If the key already exists, the value will be overwritten.
//	cache[key] = cacheEntry{
// 	createTime: time.Now(),
// 	value:      value,
func (pc *PokeCache) Write(key string, value []byte) {
	pc.mu.Lock()
	defer pc.mu.Unlock()
	pc.cache[key] = cacheEntry{
		createTime: time.Now(),
		value:      value,
	}
}

// Read returns the value associated with the key and a boolean indicating whether the key was found.
// If the key is expired, it will be deleted and the function will return false.
func (pc *PokeCache) Read(key string) ([]byte, bool) {
	pc.mu.RLock()
	defer pc.mu.RUnlock()
	entry, ok := pc.cache[key]
	if !ok {
		return nil, false
	}
	if pc.isEntryExpired(entry.createTime) {
		return nil, false
	}
	return entry.value, true
}

func (pc *PokeCache) Delete(key string) {
	pc.mu.Lock()
	defer pc.mu.Unlock()
	delete(pc.cache, key)
}

func (pc *PokeCache) Clear() {
	pc.mu.Lock()
	defer pc.mu.Unlock()
	pc.cache = make(map[string]cacheEntry)
}

func (pc *PokeCache) isEntryExpired(entryTime time.Time) bool {
	pc.mu.RLock()
	defer pc.mu.RUnlock()
	return time.Since(entryTime) > pc.maxAge
}

// cacheMaintenance is a goroutine that runs in the background and removes expired entries from the cache.
func (pc *PokeCache) cacheMaintenance() {
	defer pc.mu.Unlock()
	for {
		time.Sleep(pc.maxAge)
		pc.mu.Lock()
		for key, entry := range pc.cache {
			if time.Since(entry.createTime) > pc.maxAge {
				go pc.Delete(key)
			}
		}
		pc.mu.Unlock()
	}
}
