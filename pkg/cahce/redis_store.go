package cahce

import (
	"go1/pkg/app"
	"go1/pkg/redis"
	"time"
)

type RedisStore struct {
	RedisClient *redis.RedisClient
	PrefixKey   string
}

func (rs *RedisStore) Get(key string, clear bool) string {

	key = rs.PrefixKey + key

	value := rs.RedisClient.Get(key)

	if clear {
		rs.RedisClient.Del(key)
	}

	return value
}

func (rs *RedisStore) Set(key string, value string) bool {

	ExpireTime := time.Minute * time.Duration(5)

	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(1000)
	}

	key = rs.PrefixKey + key

	return rs.RedisClient.Set(key, value, ExpireTime)

}
