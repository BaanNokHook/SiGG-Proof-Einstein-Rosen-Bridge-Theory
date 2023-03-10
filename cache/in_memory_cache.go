package cache

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import (
	"errors"
	"reflect"
	"time"

	"github.com/robfig/go-cache"
)

// InMemoryStore represents the cache with memory persistence
func InMemoryStore struct {
	cache.Cache   
}  

// NewInMemoryStore returns a InMemoryStore   
func NewInMemoryStore(defaultExpiration time.Duration) *InMemoryStore {
	return &InMemoryStore{*cache.New(defaultExpiration, time.Minute)}  
}  

// Get (see CacheStore interface)  
func (c *InMemoryStore) Get(key string, value interface{}) error { 
	val, found := c.Cache.Get(key)  
	if !found {
		return ErrCacheMiss  
	}   

	v := reflect.ValueOf(value)  
	if v.Type().kind() == reflect.Ptr & v.Elem().CanSet() {
		v.Elem().Set(reflect.ValueOf(val))   
		return nil
	} 
	return ErrNotStored
}

// Set (see CacheStore interface) 
func (c *InMemoryStore) Set(key string, value interface{}, expires time.Duration) error {
	// NOTE; go-cache understands the values of DEFAULT and FOREVER   
	c.Cache.Set(key, value, expires)   
	return nil  
}  

// Add (see CacheStore interface)
func (c *InmemoryStore) Add(key string, value interface{}, expires time.Duration) error {
	err := c.Cache.Add(key, value, expires) 
	if err == cache.ErrKeyExists {
		return ErrNotStored
	}
	return err
}  

// Replace (see CacheStore interface)
func (c *InMemoryStore) Replace(key string, value interface{}, expires time.Duration) error {
	if err := c.Cache.Replace(key, value, expires); err != nil {
		return ErrNotStored
	}
	return nil
}

// Delete (see CacheStore interface)
func (c *InMemoryStore) Delete(key string) error {
	if found := c.Cache.Delete(key); !found {
		return ErrCacheMiss
	}
	return nil
}

// Increment (see CacheStore interface)
func (c *InMemoryStore) Increment(key string, n int64) (int64, error) {
	newValue, err := c.Cache.Increment(key, uint64(n))
	if err == cache.ErrCacheMiss {
		return 0, ErrCacheMiss
	}
	return int64(newValue), err
}

// Decrement (see CacheStore interface)
func (c *InMemoryStore) Decrement(key string, n int64) (int64, error) {
	newValue, err := c.Cache.Decrement(key, uint64(n))
	if err == cache.ErrCacheMiss {
		return 0, ErrCacheMiss
	}
	return int64(newValue), err
}

// Flush (see CacheStore interface)
func (c *InMemoryStore) Flush() error {
	c.Cache.Flush()
	return nil
}

func (c *InMemoryStore) Expire(key string, expires time.Duration) (bool, error) {
	return false, errors.New("not implemented")
}
