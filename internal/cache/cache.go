package cache

import (
	"net/url"
)

type Cache interface {
	Put(key string, value interface{})
	PutAll(map[string]interface{})
	Get(key string) interface{}
	GetAll(keys []string) map[string]interface{}
	Clean(key string)
	CleanAll()
}

func GetCachingMechanism(connStr string) Cache {
	uri, _ := url.Parse(connStr)

	switch uri.Scheme {
	case "redis":
		return GetRedisCache(connStr)
	default:
		return GetLocalCache()
	}
}
