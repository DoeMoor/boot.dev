package pokecache

import (
	"testing"
	"time"
)

func TestPokeCache_WriteAndRead(t *testing.T) {
	cache := GetCache()
	cache.Clear() // Clear the cache before starting tests

	// Test writing and reading a value
	key := "testKey"
	value := []byte("testValue")
	cache.Write(key, value)

	readValue, found := cache.Read(key)
	if !found {
		t.Errorf("Expected key %s to be found", key)
	}
	if string(readValue) != string(value) {
		t.Errorf("Expected value %s, got %s", value, readValue)
	}
}

func TestPokeCache_ExpiredEntry(t *testing.T) {
	cache := GetCache()
	cache.Clear() // Clear the cache before starting tests

	// Set maxAge to 1 second for quicker testing
	cache.SetMaxAge(1)
	key := "expiringKey"
	value := []byte("expiringValue")
	cache.Write(key, value)

	// Wait for the entry to expire
	time.Sleep(2 * time.Second)

	_, found := cache.Read(key)
	if found {
		t.Errorf("Expected key %s to be expired and not found", key)
	}
}

func TestPokeCache_Delete(t *testing.T) {
	cache := GetCache()
	cache.Clear() // Clear the cache before starting tests

	key := "deleteKey"
	value := []byte("deleteValue")
	cache.Write(key, value)

	// Ensure the key is present before deletion
	_, found := cache.Read(key)
	if !found {
		t.Errorf("Expected key %s to be found before deletion", key)
	}

	// Delete the key and check again
	cache.Delete(key)
	_, found = cache.Read(key)
	if found {
		t.Errorf("Expected key %s to be deleted", key)
	}
}

func TestPokeCache_Clear(t *testing.T) {
	cache := GetCache()
	cache.Clear() // Clear the cache before starting tests

	cache.Write("key1", []byte("value1"))
	cache.Write("key2", []byte("value2"))

	cache.Clear()

	_, found1 := cache.Read("key1")
	_, found2 := cache.Read("key2")

	if found1 || found2 {
		t.Errorf("Expected cache to be cleared")
	}
}

func TestPokeCache_Singleton(t *testing.T) {
	cache1 := GetCache()
	cache2 := GetCache()

	if cache1 != cache2 {
		t.Errorf("Expected GetCache to return the same instance")
	}
}
