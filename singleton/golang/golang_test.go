package singleton

import (
	"fmt"
	"testing"
	"time"
)

func getSingleton1(t ton) {
	t1 := GetSingleton1()
	t.Time = t1.Time
	fmt.Println(t)
}
func TestSingleton1(t *testing.T) {
	var t1, t2, t3 ton
	go getSingleton1(t1)
	time.Sleep(1e9)
	go getSingleton1(t2)
	time.Sleep(1e9)
	go getSingleton1(t3)
	time.Sleep(3e9)
	t.Log(t1, t2, t3)
}

func getSingleton2(t ton) {
	t1 := GetSingleton2()
	t.Time = t1.Time
	fmt.Println(t)
}
func TestSingleton2(t *testing.T) {
	var t1, t2, t3 ton
	go getSingleton2(t1)
	time.Sleep(1e9)
	go getSingleton2(t2)
	time.Sleep(1e9)
	go getSingleton2(t3)
	time.Sleep(3e9)
	t.Log(t1, t2, t3)
}

func getSingleton3(t ton) {
	t1 := GetSingleton3()
	t.Time = t1.Time
	fmt.Println(t)
}
func TestSingleton3(t *testing.T) {
	var t1, t2, t3 ton
	go getSingleton3(t1)
	time.Sleep(1e9)
	go getSingleton3(t2)
	time.Sleep(1e9)
	go getSingleton3(t3)
	time.Sleep(3e9)
	t.Log(t1, t2, t3)
}
