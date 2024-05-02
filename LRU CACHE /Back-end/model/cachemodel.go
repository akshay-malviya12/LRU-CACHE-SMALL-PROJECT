package models

import (
	"log"
	"sync"
	"time"
)

type CacheModel struct {
	Key             string    `json:"key"`
	Value           string    `json:"value"`
	CacheExpireTime time.Time `json:"cache_Expire_Time"`
}

type CacheLRU struct {
	cache    map[string]CacheModel
	capacity int
	mutex    sync.Mutex
}

// NewCacheLRU creates a new cache with the specified capacity.
// capacity is the maximum number of items that can be stored in the LRUcache.
func NewCacheLRU(capacity int) *CacheLRU {
	return &CacheLRU{
		capacity: capacity,
		cache:    make(map[string]CacheModel),
	}
}

// Set adds a new item to the LRU cache.
func (c *CacheLRU) Set(key string, value string, cacheExpireTime time.Time) (bool, string) {
	// Lock  locks the mutex to ensure thread-safe access to the cache.
	c.mutex.Lock()
	// defer unlock the mutex after the function is executed.
	defer c.mutex.Unlock()

	// if the cache is full, remove the least recently used item.
	log.Println("len(c.cache) >= c.capacity", len(c.cache) >= c.capacity)
	if len(c.cache) >= c.capacity {
		// delete the least recently used item from the cache.
		c.DelelteOldest()
		return false, "Failed to set cache :cache is full."
	}
	// Add the new item to the cache.
	c.cache[key] = CacheModel{Value: value, CacheExpireTime: cacheExpireTime}
	// return true to indicate that the item was added to the cache.
	if _, ok := c.cache[key]; !ok {
		return false, "failure not added in LRU cache."
	}
	return true, "successfully added LRU cache ."
}

// Get retrieves an item from the LRU cache By Provided Key.
func (c *CacheLRU) Get(key string) (CacheModel, bool) {
	// Lock  locks the mutex to ensure thread-safe access to the cache.
	c.mutex.Lock()
	// defer unlock the mutex after the function is executed.
	defer c.mutex.Unlock()
	// Check if the item exists in the cache.
	value, ok := c.cache[key]
	value.Key = key
	// return the item if it exists in the cache.
	if !ok || value.CacheExpireTime.Before(time.Now()) {
		delete(c.cache, key)
		return value, false
	}
	return value, true
}

// Delete oldest key from the LRU cache.
func (c *CacheLRU) DelelteOldest() {
	var oldestKey string
	oldestKeyExpireTime := time.Now()
	for k, v := range c.cache {
		// if the item's expire time is before the current time, delete it.
		if v.CacheExpireTime.Before(oldestKeyExpireTime) {
			oldestKeyExpireTime = v.CacheExpireTime
			oldestKey = k
		}
	}
	delete(c.cache, oldestKey)
}

// Delete cache by user key directly from the LRU cache.
func (c *CacheLRU) CacheDeleteByUserKeyDirectly(key string) bool {
	// delete the key from the cache
	//c.GetAllKeys()
	delete(c.cache, key)
	return true
}


