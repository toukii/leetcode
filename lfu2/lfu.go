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

func (c *LFUCache) Resize(size int) {
	c.Lock()
	defer c.Unlock()
	c.size = size
}

func (c *LFUCache) Insize(dlk int) {
	c.Lock()
	defer c.Unlock()
	c.size += dlk
}

func (c *LFUCache) Desize(dlk int) {
	c.Lock()
	defer c.Unlock()
	rem_c := c.size - len(c.v)
	if rem_c < dlk {
		rem_c = dlk - rem_c
		if rem_c >= len(c.v) {
			rem_c = len(c.v) - 1
		}
		for i := 0; i < rem_c; i++ {
			last := c.last
			c.remove()
			c.remv[last.Key] = last
		}
	}
	c.size -= dlk
	if c.size < 1 {
		c.size = 1
	}
}

func (c *LFUCache) Vals() []*LFU {
	c.RLock()
	defer c.RUnlock()
	ret := make([]*LFU, 0, c.size)
	for lfu := c.latest; lfu != nil; lfu = lfu.next {
		ret = append(ret, lfu)
	}
	return ret
}

func (c *LFUCache) Flush() {
	c.Lock()
	defer c.Unlock()
	for _, it := range c.v {
		it.N = 1
	}
	for _, it := range c.remv {
		it.N = 1
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

func (c *LFUCache) Attach(key string) (cur *LFU) {
	c.RLock()
	defer c.RUnlock()
	cur, exist := c.v[key]
	if exist {
		return cur
	}
	return nil
}

func (c *LFUCache) WhistPut(key string, v interface{}) (cur *LFU) {
	c.Lock()
	defer c.Unlock()
	cur, exist := c.v[key]
	// fmt.Println(exist, key, cur, c.v)
	if !exist {
		cur_remv, ok := c.remv[key]
		if ok {
			// resize
			if len(c.v) < c.size {
				cur = cur_remv
				// cur.N++
				cur.V = v
				delete(c.remv, key)
				goto ReSize_ADD
			}
			c.revoke(key)
			return cur_remv
		}
		// remove the last node
		if len(c.v) >= c.size {
			c.remove()
		}
		cur = NewLFU(key, v)
	ReSize_ADD:
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
	cur.V = v
	// cur.N++
	if cur == c.latest {
		return
	}
	c.move(cur, cur.pre)
	return
}

func (c *LFUCache) Get(key string) (cur *LFU) {
	c.Lock()
	defer c.Unlock()
	cur, exist := c.v[key]
	if !exist {
		out_cur, out_exist := c.remv[key]
		if out_exist {
			c.revoke(key)
			return out_cur
		}
		return nil
	}
	cur.N++
	if cur == c.latest {
		return
	}
	c.move(cur, cur.pre)
	return
}

func (c *LFUCache) Set(key string, v interface{}) (cur *LFU) {
	c.Lock()
	defer c.Unlock()
	cur, exist := c.v[key]
	// fmt.Println(exist, key, cur, c.v)
	if !exist {
		cur_remv, ok := c.remv[key]
		if ok {
			// resize
			if len(c.v) < c.size {
				cur = cur_remv
				cur.N++
				cur.V = v
				delete(c.remv, key)
				goto ReSize_ADD
			}
			c.revoke(key)
			return cur_remv
		}
		// remove the last node
		if len(c.v) >= c.size {
			c.remove()
		}
		cur = NewLFU(key, v)
	ReSize_ADD:
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
	cur.V = v
	cur.N++
	if cur == c.latest {
		return
	}
	c.move(cur, cur.pre)
	return
}

// remove the last node from dequeue to remv
func (c *LFUCache) remove() {
	if len(c.v) < 1 {
		return
	}
	remv := c.last
	c.last = remv.pre
	remv.pre = nil
	c.last.next = nil
	c.remv[remv.Key] = remv
	delete(c.v, remv.Key)
}

// location, from is the node before n
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
	c.v[key] = cur
}

// cur is inside of the dequeue, from is the node before cur
func (c *LFUCache) move(cur, from *LFU) {
	if cur == c.latest || from == nil {
		return
	}
	dst := c.loc(cur.N, from)

	// location is currect now
	if dst == from {
		return
	}
	c.out(cur)
	c.in(dst, cur)
	return
}

// insert cur after dst
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

// pick cur outof the dequeue
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
	cur.pre = nil
}
