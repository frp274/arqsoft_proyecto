package cache

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	log "github.com/sirupsen/logrus"
)

// CacheItem representa un item en la cache local
type CacheItem struct {
	Data      []byte
	ExpiresAt time.Time
}

// LocalCache es la cache local en memoria (primera capa)
type LocalCache struct {
	items map[string]*CacheItem
	mu    sync.RWMutex
	ttl   time.Duration
}

var (
	localCache       *LocalCache
	memcachedClient  *memcache.Client
	defaultTTL       = 5 * time.Minute
	memcachedEnabled = true
)

// InitCache initializes both cache layers
func InitCache() {
	// Initialize local cache (CCache-like implementation)
	localCache = &LocalCache{
		items: make(map[string]*CacheItem),
		ttl:   defaultTTL,
	}

	// Start cleanup goroutine for local cache
	go localCache.cleanup()

	// Initialize Memcached (distributed cache)
	memcachedHost := getEnv("MEMCACHED_HOST", "memcached")
	memcachedPort := getEnv("MEMCACHED_PORT", "11211")
	memcachedAddr := fmt.Sprintf("%s:%s", memcachedHost, memcachedPort)

	memcachedClient = memcache.New(memcachedAddr)
	memcachedClient.Timeout = 2 * time.Second

	// Test Memcached connection
	if err := memcachedClient.Ping(); err != nil {
		log.Warnf("Memcached not available at %s: %v. Using only local cache.", memcachedAddr, err)
		memcachedEnabled = false
	} else {
		log.Infof("Memcached connected at %s", memcachedAddr)
	}

	log.Info("Cache layers initialized (Local + Memcached)")
}

// Get retrieves from cache (local first, then distributed)
func Get(key string) ([]byte, bool) {
	// Try local cache first (L1)
	if data, found := localCache.Get(key); found {
		log.Debugf("Cache HIT (local): %s", key)
		return data, true
	}

	// Try Memcached (L2)
	if memcachedEnabled {
		if item, err := memcachedClient.Get(key); err == nil {
			log.Debugf("Cache HIT (memcached): %s", key)
			// Store in local cache for faster access next time
			localCache.Set(key, item.Value)
			return item.Value, true
		}
	}

	log.Debugf("Cache MISS: %s", key)
	return nil, false
}

// Set stores in both cache layers
func Set(key string, value []byte) {
	// Store in local cache (L1)
	localCache.Set(key, value)

	// Store in Memcached (L2)
	if memcachedEnabled {
		item := &memcache.Item{
			Key:        key,
			Value:      value,
			Expiration: int32(defaultTTL.Seconds()),
		}
		if err := memcachedClient.Set(item); err != nil {
			log.Warnf("Failed to set in Memcached: %v", err)
		}
	}
}

// Delete removes from both cache layers
func Delete(key string) {
	localCache.Delete(key)

	if memcachedEnabled {
		if err := memcachedClient.Delete(key); err != nil && err != memcache.ErrCacheMiss {
			log.Warnf("Failed to delete from Memcached: %v", err)
		}
	}
}

// GetStats returns cache statistics
func GetStats() map[string]interface{} {
	stats := make(map[string]interface{})

	// Local cache stats
	localCache.mu.RLock()
	stats["local_items"] = len(localCache.items)
	localCache.mu.RUnlock()

	// Memcached status
	stats["memcached_enabled"] = memcachedEnabled
	if memcachedEnabled {
		// Test ping to check if memcached is alive
		if err := memcachedClient.Ping(); err == nil {
			stats["memcached_status"] = "connected"
		} else {
			stats["memcached_status"] = "disconnected"
		}
	}

	return stats
}

// LocalCache methods

func (lc *LocalCache) Get(key string) ([]byte, bool) {
	lc.mu.RLock()
	defer lc.mu.RUnlock()

	item, found := lc.items[key]
	if !found {
		return nil, false
	}

	// Check expiration
	if time.Now().After(item.ExpiresAt) {
		return nil, false
	}

	return item.Data, true
}

func (lc *LocalCache) Set(key string, value []byte) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	lc.items[key] = &CacheItem{
		Data:      value,
		ExpiresAt: time.Now().Add(lc.ttl),
	}
}

func (lc *LocalCache) Delete(key string) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	delete(lc.items, key)
}

func (lc *LocalCache) cleanup() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		lc.mu.Lock()
		now := time.Now()
		for key, item := range lc.items {
			if now.After(item.ExpiresAt) {
				delete(lc.items, key)
			}
		}
		lc.mu.Unlock()
	}
}

// Helper functions

func GetJSON(key string, v interface{}) bool {
	data, found := Get(key)
	if !found {
		return false
	}

	if err := json.Unmarshal(data, v); err != nil {
		log.Errorf("Failed to unmarshal cache data: %v", err)
		return false
	}

	return true
}

func SetJSON(key string, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	Set(key, data)
	return nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
