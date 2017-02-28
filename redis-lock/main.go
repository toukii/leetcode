package main

import (
	"fmt"
	redis "gopkg.in/redis.v5"
	"time"
)

type LockRedis struct {
	*redis.Client
}

func ExampleNewClient() *redis.Client{
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "dvctsGH24VF5D3ccMk2FUYOCRozolDxE", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
	return client
}

func ExampleClient(client *redis.Client) {
	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exists
}

func main() {
	client:=ExampleNewClient()
	defer client.Close()
	ExampleClient(client)
	lock:=LockRedis{client}
	now:=time.Now().UnixNano()/1e6
	fmt.Println(lock.lock("lock",int64(now)))
	//fmt.Println(lock.lock("lock",now))
	//lock.unlock("lock")
	//fmt.Println(lock.lock("lock",now))
}

func (l LockRedis)lock(k string, v int64)bool  {
	for i := 0; i < 10; i++ {

		// 初次尝试获得锁
		t1:=l.SetNX(k,v,0)
		if t1.Val() { // 获得锁
			return true
		}
		fmt.Println("setnx-",t1)

		// 当前值
		t2v:=l.Get(k)
		fmt.Println("get-",t2v)
		t2,err2:= t2v.Int64()
		if err2 == redis.Nil{
			t222v:=l.GetSet(k,v)
			_,err222:=t222v.Int64()
			if err222==redis.Nil {
				return true
			}
			fmt.Println("getset-",t222v)
		}
		if v-t2>10000 { // 检查超时
			fmt.Println("timeout")
			// 尝试获得锁
			t22v:=l.GetSet(k,v)
			fmt.Println("getset-",t22v)
			t22,err22:=t22v.Int64()
			if err22 == redis.Nil || t2 == t22 { // 获得锁
				return true
			}
		}
		time.Sleep(20)
	}
  	return false
}

func (l LockRedis) unlock(k string)  {
	l.Del(k)
}