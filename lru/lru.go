package lru

import (
	"fmt"
)

type LRUCache struct {
	v         map[string]*LRU
	size      int
	lru_queue *LRU
}

func NewLRUCache(size int) *LRUCache {
	return &LRUCache{
		v:         make(map[string]*LRU),
		size:      size,
		lru_queue: nil,
	}
}

type LRU struct {
	pre, next *LRU
	key       string
	v         interface{}
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
	if c.lru_queue != nil {
		return c.lru_queue
	}
	return nil
}

func (c *LRUCache) Last() *LRU {
	if c.lru_queue != nil {
		return c.lru_queue.pre
	}
	return nil
}

func (c *LRUCache) Set(key string, v interface{}) {
	_new_lru := newLRU(key, v)
	if len(c.v) <= 0 {
		c.v[key] = _new_lru
		c.lru_queue = _new_lru
		c.lru_queue.pre = _new_lru
		c.lru_queue.next = _new_lru
		return
	}
	_, ok := c.v[key]
	if ok {
		c.v[key] = _new_lru
		c.Get(key)
		return
	}
	if len(c.v) >= c.size {
		last := c.Last()
		delete(c.v, last.key)
		c.lru_queue.pre = last.pre
		c.lru_queue.pre.next = c.lru_queue
	}
	c.v[key] = _new_lru
	_new_lru.next = c.lru_queue
	_new_lru.pre = c.lru_queue.pre
	c.lru_queue.pre = _new_lru
	c.lru_queue = _new_lru
}

func (c *LRUCache) Get(key string) interface{} {
	v, ok := c.v[key]
	if !ok {
		return nil
	}
	if c.lru_queue == v {
		return v.v
	}
	if c.lru_queue.pre == v {
		c.lru_queue.pre = v.pre
		v.pre.next = c.lru_queue
		v.next = c.lru_queue
		c.lru_queue = v
		return v.v
	}
	v.pre.next = v.next
	v.next.pre = v.pre
	v.pre = c.lru_queue.pre
	v.next = c.lru_queue
	c.lru_queue = v
	return v.v
}
