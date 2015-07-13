package singleton

import (
	"fmt"
	"sync"
	"time"
)

type ton struct {
	Time string
}

func (t *ton) String() string {
	return fmt.Sprintf("%s\t", t.Time)
}

var singleTon *ton

func GetSingleton1() *ton {
	if nil == singleTon {
		time.Sleep(3e9)
		singleTon = &ton{Time: time.Now().String()}
	}
	return singleTon
}

var locker chan bool

func init() {
	locker = make(chan bool, 1)
	locker <- true
}
func GetSingleton2() *ton {
	<-locker
	defer func() {
		locker <- true
	}()
	if nil == singleTon {
		time.Sleep(3e9)
		singleTon = &ton{Time: time.Now().String()}
	}
	return singleTon
}

var once sync.Once

func GetSingleton3() *ton {
	once.Do(func() {
		if nil == singleTon {
			time.Sleep(3e9)
			singleTon = &ton{Time: time.Now().String()}
		}
	})
	return singleTon
}
