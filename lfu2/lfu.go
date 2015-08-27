package lfu2

import (
	"fmt"
	"sync"
)

type LFUCache struct {
	last, latest *LFU
	size         int
	v            map[string]*LFU
	remv         map[string]*LFU
	sync.RWMutex
}

func NewLFUCache(size int) *LFUCache {
	if size <= 1 {
		panic("it makes no sense.")
	}
	return &LFUCache{
		size: size,
		v:    make(map[string]*LFU),
		remv: make(map[string]*LFU),
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

func (c *LFUCache) Get(key string) (cur *LFU) {
	cur, exist := c.v[key]
	if !exist {
		c.Lock()
		defer c.Unlock()
		out_cur, out_exist := c.remv[key]
		if out_exist {
			c.revoke(key)
			return out_cur
		}
		return nil
	}
	c.Lock()
	defer c.Unlock()
	cur.N++
	if cur == c.latest {
		return
	}
	c.moveFrom(cur, cur.pre)
	return
}

func (c *LFUCache) Set(key string, v interface{}) (cur *LFU) {
	c.Lock()
	defer c.Unlock()
	cur, exist := c.v[key]
	if !exist {
		cur_remv, ok := c.remv[key]
		if ok {
			c.revoke(key)
			return cur_remv
		}
		cur = NewLFU(key, v)
		c.v[key] = cur
		// cur will be the first node
		if len(c.v) <= 1 {
			c.last = cur
			c.latest = cur
			return
		}
		dst := c.loc(cur.N, c.last)
		c.in(dst, cur)
		return
	}
	cur.N++
	if cur == c.latest {
		return
	}
	c.moveFrom(cur, cur.pre)
	return
}

func (c *LFUCache) remove() {
	remv := c.last
	c.last = remv.pre
	remv.pre = nil
	c.last.next = nil
	c.remv[remv.Key] = remv
	delete(c.v, remv.Key)
}

// location
func (c *LFUCache) loc(n int, from *LFU) *LFU {
	var tmp *LFU
	for tmp = from; tmp != nil && tmp.N <= n; tmp = tmp.pre {
	}
	return tmp
}

// cur is outside of the dequeue
func (c *LFUCache) revoke(key string) {
	cur, ok := c.remv[key]
	if !ok {
		return
	}
	c.remove()
	delete(c.remv, key)
	cur.N++
	dst := c.loc(cur.N, c.last)
	c.in(dst, cur)
}

// cur is inside of the dequeue, from is the node before cur
func (c *LFUCache) moveFrom(cur, from *LFU) {
	if cur == c.latest || from == nil {
		return
	}
	dst := c.loc(cur.N, from)
	// location is currect now
	if dst == from.pre {
		return
	}
	c.out(cur)
	c.in(dst, cur)
	return
}

func (c *LFUCache) in(dst, cur *LFU) {
	// cur will be the latest node
	if dst == nil {
		cur.next = c.latest
		c.latest.pre = cur
		c.latest = cur
		return
	}
	// dst is the last node, then cur will be the last node
	if c.last == dst {
		cur.pre = c.last
		c.last.next = cur
		c.last = cur
		return
	}
	// insert cur after dst
	cur.next = dst.next
	cur.pre = dst
	cur.next.pre = cur
	dst.next = cur
}

func (c *LFUCache) out(cur *LFU) {
	// if cur == nil || cur.pre == nil {
	// 	panic("cur must not be nil")
	// }
	if c.last != cur {
		cur.next.pre = cur.pre
	} else {
		c.last = cur.pre
		c.last.next = nil
	}
	cur.pre.next = cur.next
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
