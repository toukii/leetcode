package lru

import (
	"testing"
)

func assert(t *testing.T, k1, k2 string) {
	if k1 != k2 {
		t.Errorf("want %s, but %s.", k2, k1)
	}
}

func TestLRUSet(t *testing.T) {
	cache := NewLRUCache(2)
	cache.Set("1", 1)
	assert(t, cache.Last().Key, "1")
	assert(t, cache.Latest().Key, "1")

	cache.Set("2", "two")
	assert(t, cache.Last().Key, "1")
	assert(t, cache.Latest().Key, "2")

	cache.Set("3", "three")
	assert(t, cache.Last().Key, "2")
	assert(t, cache.Latest().Key, "3")
}

func TestLRUSetGet(t *testing.T) {
	cache := NewLRUCache(2)
	cache.Set("1", 1)
	assert(t, cache.Last().Key, "1")
	assert(t, cache.Latest().Key, "1")

	cache.Set("2", "two")
	assert(t, cache.Last().Key, "1")
	assert(t, cache.Latest().Key, "2")

	cache.Get("1")
	assert(t, cache.Last().Key, "2")
	assert(t, cache.Latest().Key, "1")

	cache.Set("3", "three")
	assert(t, cache.Last().Key, "1")
	assert(t, cache.Latest().Key, "3")

	cache.Get("1")
	assert(t, cache.Last().Key, "3")
	assert(t, cache.Latest().Key, "1")
}
