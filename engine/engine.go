package engine

import (
	"sync"
	"time"
)

var (
	store = make(map[string]KVValue)
	mu    sync.Mutex
)

type KVValue struct {
	Value  string
	Expiry int64
}

func setKeyValue(key, value string, expireTimeMillis int64) {
	mu.Lock()
	defer mu.Unlock()

	store[key] = KVValue{
		Value:  value,
		Expiry: expireTimeMillis,
	}
}

func getValue(key string) string {
	mu.Lock()
	defer mu.Unlock()

	if kv, exists := store[key]; exists {
		if kv.Expiry == 0 || time.Now().UnixMilli() < kv.Expiry {
			return kv.Value
		}
		delete(store, key)
	}
	return "(nil)"
}
