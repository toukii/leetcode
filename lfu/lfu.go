package lfu

import (
	"fmt"
	"sync"
)

type LFUCache struct {
	last, latest *LFU
	size         int
	v            map[string]*LFU
	outv         map[string]*LFU
	sync.RWMutex
}

func NewLFUCache(size int) *LFUCache {
	if size <= 1 {
		panic("it makes no sense.")
	}
	return &LFUCache{
		size: size,
		v:    make(map[string]*LFU),
		outv: make(map[string]*LFU),
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
		c.Lock()
		defer c.Unlock()
		out_cur, out_exist := c.outv[key]
		if out_exist {
			out_cur.N++
			delete(c.outv, key)
			c.v[key] = out_cur
			out_cur.pre = c.last
			out_cur.next = nil
			c.last.next = out_cur
			c.last = out_cur
			// del pre
			pre := c.last.pre
			pre.pre.next = pre.next
			pre.next.pre = pre.pre
			delete(c.v, pre.Key)
			c.outv[pre.Key] = pre
			c.moveForward(key)
			return out_cur
		}
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
	if cur.N >= cur.pre.N {
		c.moveForward(key)
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
			if len(c.v) >= c.size {
				// delete the last node
				// c.last.pre = nil
				// c.last.pre.next = nil
				delete(c.v, c.last.Key)
				c.outv[c.last.Key] = c.last
				c.last = c.last.pre
				c.last.next = nil
			}
			cur.pre = c.last
			c.last.next = cur
		}
		// new node is the last node
		c.last = cur
		c.v[key] = cur
	} else {
		// update cache key
		cur.V = v
		cur.N++
	}
	if cur.pre != nil && cur.pre.N <= cur.N {
		c.moveForward(key)
	}
	return
}

func (c *LFUCache) moveForward(key string) {
	cur, ok := c.v[key]
	if !ok {
		panic("cur key must be not nil")
	}
	// only one node
	if len(c.v) <= 1 {
		return
	}
	if c.latest == cur {
		return
	}
	if cur.pre != nil {
		if cur.pre.N > cur.N {
			return
		}
		if c.last == cur {
			c.last = cur.pre
		}
		cur.pre.next = cur.next
		cur.pre = nil
	}
	// cur is the last node
	if cur.next != nil {
		cur.next.pre = cur.pre
	}
	var tmp *LFU
	for tmp = cur.pre; tmp != nil && tmp.N <= cur.N; tmp = tmp.pre {
	}
	// cur will be the latest node
	if tmp == nil {
		c.latest.pre = cur
		cur.next = c.latest
		c.latest = cur
		return
	}
	// insert cur after tmp
	if tmp == cur.pre {
		return
	}
	cur.pre.next = cur.next

	cur.next = tmp.next
	cur.pre = tmp
	tmp.next.pre = cur
	tmp.next = cur
}
