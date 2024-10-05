package locations_map

import (
	"net/http"
	"sync"
	"time"
	"encoding/json"
)

type locations struct { 
	// Count    int64       `json:"count"`
	Next      *string `json:"next"`
	Previous  *string `json:"previous"`
	Results   []location `json:"results"`
	timeStamp time.Time
}

type location struct {
	Name string `json:"name"`
	URL string `json:"url"`
}


func (lc *locations) newLocations( rawResp *http.Response) error {
	decoder := json.NewDecoder(rawResp.Body)
	err := decoder.Decode(lc)
	if err != nil {
		return err
	}
	lc.timeStamp = time.Now()
	return nil
}

type Cache struct {
	cache map[string]locations
	mu 	sync.RWMutex
}

var LocationCache *Cache
var lock sync.Mutex

func GetCache() *Cache {
	if LocationCache == nil {
		lock.Lock()
		defer lock.Unlock()
		if LocationCache == nil {
			LocationCache = &Cache{
				cache: make(map[string]locations),
			}
		}
	}
	return LocationCache
}

func (locationCache *Cache) write(urlKey string, locations locations) {
	locationCache.mu.Lock()
	defer locationCache.mu.Unlock()

	locationCache.cache[urlKey] = locations
}

func (locationCache *Cache) read(urlKey string) (locations, bool) {
	locationCache.mu.RLock()
	defer locationCache.mu.RUnlock()

	cachedLocation, ok := locationCache.cache[urlKey]
	if !ok {
		return locations{}, false
	}
	
	return cachedLocation, ok
}

func (locationCache *Cache) dell(urlKey string) {
	locationCache.mu.Lock()
	defer locationCache.mu.Unlock()
	delete(locationCache.cache, urlKey)
}

func (locationCache *Cache) clear() {
	locationCache.mu.Lock()
	defer locationCache.mu.Unlock()

	locationCache.cache = make(map[string]locations)
}
