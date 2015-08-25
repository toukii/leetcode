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
	assert(t, cache.Last().key, "1")
	assert(t, cache.Latest().key, "1")

	cache.Set("2", "two")
	assert(t, cache.Last().key, "1")
	assert(t, cache.Latest().key, "2")

	cache.Set("3", "three")
	assert(t, cache.Last().key, "2")
	assert(t, cache.Latest().key, "3")
}

func TestLRUSetGet(t *testing.T) {
	cache := NewLRUCache(2)
	cache.Set("1", 1)
	assert(t, cache.Last().key, "1")
	assert(t, cache.Latest().key, "1")

	cache.Set("2", "two")
	assert(t, cache.Last().key, "1")
	assert(t, cache.Latest().key, "2")

	cache.Get("1")
	assert(t, cache.Last().key, "2")
	assert(t, cache.Latest().key, "1")

	cache.Set("3", "three")
	assert(t, cache.Last().key, "1")
	assert(t, cache.Latest().key, "3")

	cache.Get("1")
	assert(t, cache.Last().key, "3")
	assert(t, cache.Latest().key, "1")
}
