package ratelimit

import (
	"fmt"
	"sync"
	"time"
)

type LimitType int

const (
	Per60secLimit LimitType = iota
	Per30secLimit
	Per10secLimit
)

type Rate struct {
	Rate  float32
	LType LimitType

	head      int // 当前head游标
	tail      int // 当前tail游标
	capacity  int
	timestamp []int32
	last      int32

	sync.RWMutex
}

func (r *Rate) Init() {
	switch r.LType {
	case Per60secLimit:
		r.capacity = int(r.Rate * 6.0)
		// r.last = 60
		r.last = 6
	case Per30secLimit:
		r.capacity = int(r.Rate * 3.0)
		// r.last = 30
		r.last = 3
	default:
		r.capacity = int(r.Rate * 1.0)
		// r.last = 10
		r.last = 1
	}
	r.timestamp = make([]int32, r.capacity)
	go func() {
		r.tick()
	}()
}

func (r *Rate) tick() {
	ticker := time.NewTicker(time.Second)
	now := time.Now().Unix()
	for {
		now++
		<-ticker.C
		if r.Limit() <= 0 {
			continue
		}
		cur := r.tail
		r.Lock()
		for i := r.tail; (i+r.capacity)%r.capacity < r.head; i++ {
			c := (i + cur) % r.capacity
			t := r.timestamp[c]
			if t > 0 && int32(now)-t >= r.last {
				r.timestamp[c] = 0
				r.tail = (r.tail + 1) % r.capacity
			}
		}
		fmt.Printf("head:%d ,tail:%d, left:%d ,%+v\n", r.head, r.tail, r.Limit(), r.timestamp)
		r.Unlock()
	}
}

func (r *Rate) Req() (ok bool, err error) {
	if r.Limit() <= 0 {
		return false, fmt.Errorf("no limit left.")
	}
	r.Accept()
	return true, nil
}

func (r *Rate) Accept() {
	r.Lock()
	defer r.Unlock()
	r.timestamp[r.head] = int32(time.Now().Unix())
	r.head = (r.head + 1) % r.capacity
}

func (r *Rate) Limit() int {
	if r.tail == r.head {
		return r.capacity
	}
	return (r.tail + r.capacity - r.head) % r.capacity
}
