package cahce

import (
	"go1/pkg/redis"
	"sync"
)

type Cache struct {
	Store Store
}

var once sync.Once
var internalCache *Cache

func NewCache() *Cache {
	once.Do(func() {
		internalCache = &Cache{
			Store: &RedisStore{
				RedisClient: redis.Redis,
				PrefixKey:   "redis:cache",
			},
		}
	})

	return internalCache
}

func (c *Cache) GetCache(key string, clear bool) (value string) {
	val := c.Store.Get(key, clear)
	return val
}

func (c *Cache) SetCache(key string, value string) bool {
	return c.Store.Set(key, value)
}
