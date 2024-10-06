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

//	createTime time.Time
//	value      []byte // response body
type cacheEntry struct {
	createTime time.Time
	value      []byte // response body
}

var cache *PokeCache
var lock sync.Mutex

func GetCache() *PokeCache {
	if cache == nil {
		lock.Lock()
		defer lock.Unlock()
		if cache == nil {
			cache = &PokeCache{
				cache:  make(map[string]cacheEntry),
				maxAge: 5 * time.Second,
			}
		}
	}
	return cache
}

func (pc *PokeCache) Write(key string, value []byte) {
	pc.mu.Lock()
	defer pc.mu.Unlock()
	pc.cache[key] = cacheEntry{
		createTime: time.Now(),
		value:      value,
	}
}

func (pc *PokeCache) Read(key string) ([]byte, bool) {
	pc.mu.RLock()
	defer pc.mu.RUnlock()
	entry, ok := pc.cache[key]
	if !ok {
		return nil, false
	}
	if pc.isEntryExpired(entry.createTime) {
		pc.mu.RUnlock()
		pc.Delete(key)
		pc.mu.RLock()
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
