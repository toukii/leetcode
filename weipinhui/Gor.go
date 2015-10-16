package main

import (
	"fmt"
	"time"
)

var (
	chanMap map[string]chan bool
)

func init() {
	chanMap = make(map[string]chan bool)
	chanMap["A"] = make(chan bool, 1)
	chanMap["B"] = make(chan bool, 1)
	chanMap["C"] = make(chan bool, 1)
}

func nextName(name string) string {
	ret := ""
	switch name {
	case "A":
		ret = "B"
	case "B":
		ret = "C"
	case "C":
		ret = "A"
	}
	return ret
}

func goruntine(name string, chanMap map[string]chan bool) {
	for i := 0; i < 10; i++ {
		<-chanMap[name]
		fmt.Print(name)
		chanMap[nextName(name)] <- true
	}
}

func main() {
	chanMap["A"] <- true
	go goruntine("A", chanMap)
	go goruntine("B", chanMap)
	go goruntine("C", chanMap)
	time.Sleep(1e9)
}
