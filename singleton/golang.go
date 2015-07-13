package singleton

import (
	"sync"
	"time"
)

type ton struct {
	Time string
}

var singleTon *ton

func GetSingleton1() *ton {
	if nil == singleTon {
		singleTon = &ton{}
	}
	return singleTon
}

var locker chan bool

func init() {
	locker = make(chan bool)
	locker <- true
}
func GetSingleton2() *ton {
	<-locker
	defer func() {
		locker <- true
	}()
	if nil == singleTon {
		singleTon = &ton{}
	}
	return singleTon
}

func GetSingleton3() *ton {
	sync.Once(func() {
		if nil == singleTon {
			singleTon = &ton{}
		}
	})
	return singleTon
}
