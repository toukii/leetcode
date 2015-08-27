package lfu_

import (
	"fmt"
	"sync"
)

type LFUCache struct {
	last, latest *LFU
	size         int
	v            map[string]*LFU
	sync.RWMutex
}

func NewLFUCache(size int) *LFUCache {
	return &LFUCache{
		size: size,
		v:    make(map[string]*LFU),
	}
}

func (c *LFUCache) Display() {
	c.RLock()
	defer c.RUnlock()
	first := true
	i := 0
	for lfu := c.latest; lfu != nil; lfu = lfu.next {
		i++
		if i > 10 {
			break
		}
		if lfu == c.latest && !first {
			break
		}
		first = false
		fmt.Print(lfu.String(), " //// ")
	}
	fmt.Println()
}

func (c *LFUCache) Latest() *LFU {
	c.RLock()
	defer c.RUnlock()
	return c.latest
}

func (c *LFUCache) Last() *LFU {
	c.RLock()
	defer c.RUnlock()
	return c.last
}

type LFU struct {
	Key       string
	V         interface{}
	N         int
	pre, next *LFU
	sub       *subNode
}

func NewLFU(key string, v interface{}) *LFU {
	return &LFU{
		Key: key,
		V:   v,
		N:   1,
	}
}

func (l *LFU) String() string {
	return fmt.Sprintf("<%d times> %s:%v", l.N, l.Key, l.V)
}

type subNode struct {
	unix int64
}

func (c *LFUCache) Get(key string) (cur *LFU) {
	cur, exist := c.v[key]
	if !exist {
		return nil
	}
	c.Lock()
	defer c.Unlock()
	cur.N++
	// only one node
	if len(c.v) <= 1 {
		return
	}
	// first node
	if cur.pre == nil {
		return
	}
	// move forward
	if cur.N > cur.pre.N {
		// not the last node
		if cur.next != nil {
			cur.next.pre = cur.pre
		} else {
			// cur's pre will be the last
			c.last = cur.pre
		}
		cur.pre.next = cur.next

		pre := cur.pre
		// pre is the first node
		if c.latest == pre {
			// cur will be the first node
			c.latest = cur
			cur.next = pre
			pre.pre = cur
			return
		}
		pre.pre.next = cur
		pre.pre = cur
		cur.next = pre
		cur.pre = pre.pre
	}
	return
}

func (c *LFUCache) Set(key string, v interface{}) (cur *LFU) {
	c.Lock()
	defer c.Unlock()
	cur, exist := c.v[key]
	if !exist {
		cur = NewLFU(key, v)
		// cur is the first node
		if len(c.v) <= 0 {
			c.latest = cur
		} else {
			cur.pre = c.last
			c.last.next = cur
		}
		// new node is the last node
		c.last = cur
		c.v[key] = cur
		return
	}
	// update cache key
	cur.V = v
	c.Unlock()
	c.Get(key)
	c.Lock()
	return
}
