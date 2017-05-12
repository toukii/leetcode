package main

import (
	"time"
	"fmt"
)

func main() {
	file:= File{
		locs : []string{"A","B","C","D"},
		files:make(map[string][]int),
		buf: make(chan int,10),
	}
	for _,it := range file.locs{
		file.files[it]=make([]int,0,10)
	}
	go func() {
		for i := 0; i < 200; i++ {
			file.buf<-i%4+1;
			time.Sleep(1e9)
		}
	}()
	var v int
	cur:=0
	l:=len(file.locs)
	for{
		v=<-file.buf
		file.files[file.locs[cur]]=append(file.files[file.locs[cur]],v)
		if v==l{
			file.files[file.locs[cur]]=append(file.files[file.locs[cur]],<-file.buf)
		}
		cur++
		if cur>=4 {
			cur = 0
		}
		time.Sleep(5e8)
		file.Log()
	}

}

type File struct {
	locs []string
	files map[string][]int
	buf chan int
}

func (f File) Log()  {
	fmt.Println("======================")
	for _,it:=range f.locs{
		fmt.Println(it,":",f.files[it])
	}
}

