package cache

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import (
	"errors"
	"time"

	"github.com/primasio/wormhole/config"
)

var cacheStore CacheStore
var cacheType string

func InitCache() error {
	c := config.GetConfig()

	cacheType = c.GetString("cache.type")

	if cacheType == "memory" {
		cacheStore = NewInMemoryStore(time.Second)
	} else if cacheType == "redis" {

		host := c.GetString("cache.host")
		port := c.GetString("cache.port")
		password := c.GetString("cache.password")

		// use our own redis cache since the originalversion is poorly written
		cacheStore = NewRedisCache(host+":"+port, password, time.Second)
	} else {
		return errors.New("unrecoginzed cache type")
	}

	return nil
}

func GetCache() CacheStore {
	return cacheStore
}

func GetCacheType() string {
	return cacheType
}
