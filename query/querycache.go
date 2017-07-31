package query

import (
	"time"

	servers "github.com/alanfran/SteamCondenserGo"
)

// Cache queries servers for information and caches the results.
type Cache struct {
	cache  map[string]CachedInfo
	maxAge time.Duration
}

// CachedInfo stores a server response along with its time of creation.
type CachedInfo struct {
	servers.Response
	created time.Time
}

func NewCache(maxAge time.Duration) *Cache {
	return &Cache{
		cache:  make(map[string]CachedInfo),
		maxAge: maxAge,
	}
}

// Get returns a cached response, or requests a fresh one from the game server.
func (c *Cache) Get(address string) (servers.Response, error) {
	cached, ok := c.cache[address]
	if ok && time.Now().Before(cached.created.Add(c.maxAge)) {
		return cached.Response, nil
	}

	response, err := servers.QueryGoldServer(address)
	if err == nil {
		c.Put(address, response)
	}

	return response, err
}

// Put stores a response in the cache.
func (c *Cache) Put(address string, result servers.Response) error {
	c.cache[address] = CachedInfo{Response: result, created: time.Now()}
	return nil
}

// Del removes a response from the cache.
func (c *Cache) Del(address string) {
	delete(c.cache, address)
}

// Purge will delete any cached responses older than the maximum age.
func (c *Cache) Purge() {
	for k := range c.cache {
		if time.Now().After(c.cache[k].created.Add(c.maxAge)) {
			c.Del(k)
		}
	}
}
