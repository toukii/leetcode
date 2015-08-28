package lfu2

import (
	"fmt"
	"testing"
)

func assert(t *testing.T, msg, k1, k2 string) {
	if k1 != k2 {
		t.Errorf("%s want %s, but %s.", msg, k2, k1)
	}
}

func TestLFUCache(t *testing.T) {
	t.Parallel()
	cache := NewLFUCache(4)
	cache.Set("1", 1)
	cache.Display()
	// 1:1
	assert(t, "Last", cache.Last().Key, "1")
	assert(t, "Latest", cache.Latest().Key, "1")

	cache.Set("2", "two")
	cache.Display()
	// 2:1, 1:1
	assert(t, "Last", cache.Last().Key, "1")
	assert(t, "Latest", cache.Latest().Key, "2")

	cache.Get("2")
	cache.Display()
	// 2:2, 1:1
	assert(t, "Last", cache.Last().Key, "1")
	assert(t, "Latest", cache.Latest().Key, "2")

	cache.Get("1")
	cache.Display()
	// 1:2, 2:2
	assert(t, "Last", cache.Last().Key, "2")
	assert(t, "Latest", cache.Latest().Key, "1")

	cache.Set("1", "one")
	cache.Display()
	// 1:3, 2:2
	assert(t, "Last", cache.Last().Key, "2")
	assert(t, "Latest", cache.Latest().Key, "1")

	cache.Get("2")
	cache.Display()
	// 2:3, 1:3
	assert(t, "Last", cache.Last().Key, "1")
	assert(t, "Latest", cache.Latest().Key, "2")

	cache.Set("3", "three")
	cache.Display()
	// 2:3, 1:3, 3:1
	assert(t, "Last", cache.Last().Key, "3")
	assert(t, "Latest", cache.Latest().Key, "2")

	cache.Set("3", "three")
	cache.Display()
	// 2:3, 1:3, 3:2
	assert(t, "Last", cache.Last().Key, "3")
	assert(t, "Latest", cache.Latest().Key, "2")

	cache.Set("3", "three")
	cache.Display()
	// 3:3, 2:3, 1:3
	assert(t, "Last", cache.Last().Key, "1")
	assert(t, "Latest", cache.Latest().Key, "3")
}

func TestLFUCacheSize(t *testing.T) {
	t.Parallel()
	cache := NewLFUCache(2)
	cache.Set("1", 1)
	cache.Display()
	// 1:1
	assert(t, "Last", cache.Last().Key, "1")
	assert(t, "Latest", cache.Latest().Key, "1")

	cache.Set("2", "two")
	cache.Display()
	// 2:1, 1:1
	assert(t, "Last", cache.Last().Key, "1")
	assert(t, "Latest", cache.Latest().Key, "2")

	cache.Set("3", "three")
	cache.Display()
	// 3:1, 2:1   {1:1}
	assert(t, "Last", cache.Last().Key, "2")
	assert(t, "Latest", cache.Latest().Key, "3")

	cache.Get("1")
	cache.Display()
	// 1:2, 3:1   {2:1}
	assert(t, "Last", cache.Last().Key, "3")
	assert(t, "Latest", cache.Latest().Key, "1")

	cache.Set("1", "1")
	cache.Display()
	// 1:3, 3:1   {2:1}
	assert(t, "Last", cache.Last().Key, "3")
	assert(t, "Latest", cache.Latest().Key, "1")

	cache.Set("2", "two...")
	cache.Display()
	// 1:3, 2:2   {3:1}
	assert(t, "Last", cache.Last().Key, "2")
	assert(t, "Latest", cache.Latest().Key, "1")

	cache.Get("2")
	cache.Display()
	// 2:3, 1:3   {3:1}
	assert(t, "Last", cache.Last().Key, "1")
	assert(t, "Latest", cache.Latest().Key, "2")

	cache.Get("3")
	cache.Display()
	// 2:3, 3:2   {1:3}
	assert(t, "Last", cache.Last().Key, "3")
	assert(t, "Latest", cache.Latest().Key, "2")

	cache.Set("3", "3333")
	cache.Display()
	// 3:3, 2:3   {1:3}
	assert(t, "Last", cache.Last().Key, "2")
	assert(t, "Latest", cache.Latest().Key, "3")

	cache.Get("1")
	cache.Display()
	// 1:4, 3:3   {2:3}
	assert(t, "Last", cache.Last().Key, "3")
	assert(t, "Latest", cache.Latest().Key, "1")

}

func BenchmarkCache(b *testing.B) {
	cache := NewLFUCache(100)
	for i := 0; i < b.N; i++ {
		cache.Set("1", "1")
	}
}

func BenchmarkCache2(b *testing.B) {
	cache := NewLFUCache(100)
	for i := 0; i < b.N; i++ {
		cache.Set(fmt.Sprintf("%d", i%200), "1")
	}
}