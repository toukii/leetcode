package main

import (
	"fmt"
	redis "gopkg.in/redis.v5"
	"github.com/go-macaron/macaron"
)


var(
	lock LockRedis
)

func test()  {
	lock.Set("testdecr",5,-1)
	n:=30
	c:=make(chan int,n)

	for i := 0; i < n; i++ {
		go func(){
			ok:=lock.decr("testdecr")
			if ok{
				c<-1
			}else{
				c<-0
			}
		}()
	}
	count:=0
	for i:=0;i<n;i++ {
		if <-c==1{
		count++
		}
	}
	fmt.Println(count)
}

func main() {
	client:=ExampleNewClient()
	defer client.Close()
	lock=LockRedis{client}
	//test()
	//return
	m := macaron.Classic()
	m.Group("/", func() {
		m.Get("", get)
		m.Get("set", set)
		m.Combo("decr").Get(decr)
	})
	m.Run(48080)
}

func get(ctx *macaron.Context) string {
	k:=ctx.Query("k")
	return lock.Get(k).String()
}


func set(ctx *macaron.Context) string {
	k:=ctx.Query("k")
	v:=ctx.Query("v")
	fmt.Println("set ",k,":",v,"result:",lock.Set(k,v,-1).String())
	return lock.Get(k).String()
}

func decr(ctx *macaron.Context) string  {
	k:=ctx.Query("k")
	return fmt.Sprint(lock.decr(k))
}

type LockRedis struct {
	*redis.Client
}

func ExampleNewClient() *redis.Client{
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "7rOszAMIBZENBIqa7zcp2zS9fTn8OrKW", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
	return client
}


func (l LockRedis) decr(k string)bool  {
	watchErr:= l.Watch(func(tx *redis.Tx)error{
		vrs:=tx.Get(k)
		v,err:=vrs.Int64()
		if(err!=nil){
			return err
		}
		if v<=0{
			//fmt.Println(tx.Set(k,0,-1))
			return fmt.Errorf("zero")
		}
		/*dec:=tx.Decr(k)
		fmt.Println(dec,dec.Err())
		return tx.Decr(k).Err()*/
		//status :=tx.Set(k,v-1,-1)
		//return status.Err()
		status:=tx.GetSet(k,v-1)
		fmt.Println(status,status.Err())
		old,err2:=status.Int64()
		if err2!=nil{
			return err2
		}

		if old < v{
			return fmt.Errorf("God, GETSET(old value):%d; bug GET(old value):%d, SET:%d.",old,v,v-1)
		}
		if old != v{
			return fmt.Errorf("currence!")
		}

		return status.Err()
	},k)
	//fmt.Println("watch:",k,",result:",watchErr)
	return nil== watchErr
}