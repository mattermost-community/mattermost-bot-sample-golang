package cache

import (
	"context"
	"fmt"
	"net/url"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	conn *redis.Client
	ctx  context.Context
}

func GetRedisCache(connStr string) *RedisCache {
	uri, _ := url.Parse(connStr)
	password, _ := uri.User.Password()

	cch := &RedisCache{
		conn: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", uri.Hostname(), uri.Port()),
			Username: uri.User.Username(),
			Password: password,
		}),
	}

	cch.ctx = context.Background()

	return cch
}

func (rc *RedisCache) Put(key string, value interface{}) {
	if err := rc.conn.Set(rc.ctx, key, value, 0); err != nil {
		fmt.Println(err)
	}
}

func (rc *RedisCache) PutAll(entries map[string]interface{}) {
	for k, v := range entries {
		rc.Put(k, v)
	}
}

func (rc *RedisCache) Get(key string) interface{} {
	value, err := rc.conn.Get(rc.ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return value
}

func (rc *RedisCache) GetAll(keys []string) map[string]interface{} {
	entries := make(map[string]interface{})
	for _, k := range keys {
		entries[k] = rc.Get(k)
	}

	return entries
}

func (rc *RedisCache) Clean(key string) {
	if err := rc.conn.Del(rc.ctx, key); err != nil {
		fmt.Println(err)
	}
}

// CleanAll cleans the entire cache.
func (rc *RedisCache) CleanAll() {
	rc.conn.FlushDB(rc.ctx)
}
