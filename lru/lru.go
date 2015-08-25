package lru

import (
	"fmt"
)

type LRUCache struct {
	v    map[string]*LRU
	size int
	LRU
}

func NewLRUCache(size int) *LRUCache {
	return &LRUCache{
		v:    make(map[string]*LRU),
		size: size,
		LRU:  LRU{},
	}
}

func (c *LRUCache) Display() {
	first := true
	for lru := c.next; lru != nil; lru = lru.next {
		if lru == c.next && !first {
			break
		}
		first = false
		fmt.Print(lru.String(), " ")
	}
	fmt.Println()
}

type LRU struct {
	pre, next *LRU
	key       string
	v         interface{}
}

func (l *LRU) del() {
	l.next.pre = l.pre
	l.pre.next = l.next
}

func (l *LRU) String() string {
	return fmt.Sprintf("%s:%v", l.key, l.v)
}

func newLRU(k string, v interface{}) *LRU {
	return &LRU{
		key:  k,
		v:    v,
		pre:  nil,
		next: nil,
	}
}

func (c *LRUCache) Latest() *LRU {
	return c.next
}

func (c *LRUCache) Last() *LRU {
	return c.pre
}

func (c *LRUCache) Set(key string, v interface{}) {
	_new_lru := newLRU(key, v)
	if len(c.v) <= 0 {
		c.v[key] = _new_lru
		c.next = _new_lru
		c.pre = _new_lru
		_new_lru.next = _new_lru
		_new_lru.pre = _new_lru
		return
	}
	cur, ok := c.v[key]
	if ok {
		cur.del()
		c.v[key] = _new_lru
		cur = c.v[key]
		cur.pre = c.pre
		cur.next = c.next
		c.next = cur
		cur.next.pre = cur
		cur.pre.next = cur
		return
	}
	if len(c.v) >= c.size {
		last := c.Last()
		delete(c.v, last.key)
		last.del()
		if c.pre == last {
			c.pre = last.pre
		}
	}
	c.v[key] = _new_lru
	cur = c.v[key]
	cur.pre = c.pre
	cur.next = c.next
	c.next = cur
	cur.next.pre = cur
	cur.pre.next = cur
}

func (c *LRUCache) Get(key string) interface{} {
	cur, ok := c.v[key]
	if !ok {
		return nil
	}
	if c.next == cur {
		return cur.v
	}
	cur.del()
	if c.pre == cur {
		c.pre = cur.pre
	}
	cur.pre = c.pre
	cur.next = c.next
	c.next = cur
	cur.next.pre = cur
	cur.pre.next = cur
	return cur.v
}
