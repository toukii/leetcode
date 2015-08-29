// 0.00110011001100110011001100110011
package main

import (
	"fmt"
	"sync"
	"time"
)

type Center struct {
	finished chan bool
	all      int64
	sync.Mutex
	Users map[string]*User
}

type User struct {
	Name  string
	Money int64
}

func (u *User) Buy(center *Center) error {
	return center.Buy(u)
}

func (c *Center) Buy(from *User) error {
	if from.Money <= 0 || from.Money%1000 != 0 || from.Money > 2e4 {
		return fmt.Errorf("所购金额不符。")
	}
	fmt.Println(from)
	c.Lock()
	defer c.Unlock()
	if c.all < from.Money {
		return fmt.Errorf("余额不足")
	}
	if u, ok := c.Users[from.Name]; ok {
		if u.Money+from.Money > 2e4 {
			return fmt.Errorf("金额超出20000元限制。")
		}
		u.Money += from.Money
		c.Users[from.Name] = u
	} else {
		c.Users[from.Name] = from
	}
	c.all -= from.Money
	if c.all <= 0 {
		c.finished <- true
	}
	return nil
}

func (c *Center) Show() {
	for _, u := range c.Users {
		fmt.Printf("user:%s buy:%d￥\n", u.Name, u.Money)
	}
}

func main() {
	center := &Center{finished: make(chan bool), all: 2e8, Users: make(map[string]*User)}
	u1 := &User{Name: "jack", Money: 2000}
	u2 := &User{Name: "tom", Money: 4000}

	err := u1.Buy(center)
	fmt.Println(err)

	go u2.Buy(center)
	err = u2.Buy(center)
	fmt.Println(err)

	// debug code
	time.Sleep(1e9)
	// <-center.finished
	center.Show()
}
