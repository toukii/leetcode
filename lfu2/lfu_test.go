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

func TestLFUCacheResize(t *testing.T) {
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

	cache.Resize(3)
	cache.Set("1", "one")
	cache.Display()
	// 1:2, 3:1 , 2:1
	assert(t, "Last", cache.Last().Key, "2")
	assert(t, "Latest", cache.Latest().Key, "1")

	cache.Set("2", "twoo")
	cache.Display()
	// 2:2, 1:2, 3:1
	assert(t, "Last", cache.Last().Key, "3")
	assert(t, "Latest", cache.Latest().Key, "2")
}

func TestLFUCacheWhist(t *testing.T) {
	// t.Parallel()
	cache := NewLFUCache(2)
	cur := cache.Attach("1")
	// 1:1
	if cur != nil {
		t.Error(cur)
	}
	cur = cache.WhistPut("1", "one")
	assert(t, "Last", cur.Key, "1")
	assert(t, "Last", cache.Last().Key, "1")
	assert(t, "Latest", cache.Latest().Key, "1")

	cur = cache.Attach("1")
	cache.Display()
	// 1:1
	assert(t, "Last", cur.Key, "1")
	assert(t, "Last", cache.Last().Key, "1")
	assert(t, "Latest", cache.Latest().Key, "1")

	cur = cache.WhistPut("2", "two")
	cache.Display()
	// 2:1, 1:1
	assert(t, "Last", cur.Key, "2")
	assert(t, "Last", cache.Last().Key, "1")
	assert(t, "Latest", cache.Latest().Key, "2")

	cur = cache.WhistPut("2", "twoOO")
	cache.Display()
	// 2:1, 1:1
	assert(t, "Last", cur.Key, "2")
	assert(t, "Last", cache.Last().Key, "1")
	assert(t, "Latest", cache.Latest().Key, "2")
}

func TestLFUCacheDesize(t *testing.T) {
	// t.Parallel()
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

	cache.Resize(3)
	cache.Desize(1)
	cache.Set("1", "one")
	cache.Display()
	// 1:2, 3:1  {2:1}
	assert(t, "Last", cache.Last().Key, "3")
	assert(t, "Latest", cache.Latest().Key, "1")

	cache.Desize(1)
	cache.Display()
	// 1:2   {3:1,2:1}
	assert(t, "Last", cache.Last().Key, "1")
	assert(t, "Latest", cache.Latest().Key, "1")

	cache.Desize(1)
	cache.Display()
	// 1:2   {3:1,2:1}
	assert(t, "Last", cache.Last().Key, "1")
	assert(t, "Latest", cache.Latest().Key, "1")
}

type data struct {
	key string
	val interface{}
}

var dat []data
var cache *LFUCache

func init() {
	dat = make([]data, 100)
	for i := 0; i < 100; i++ {
		dat[i] = data{key: fmt.Sprintf("%d", i), val: i}
	}
	cache = NewLFUCache(100)
}

func BenchmarkCacheSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cache.Set(dat[i%100].key, dat[i%100].val)
	}
}

func BenchmarkCacheGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cache.Get(dat[i%100].key)
	}
}
