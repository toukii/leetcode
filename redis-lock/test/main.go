package main

import (
	"fmt"
	redis "gopkg.in/redis.v5"
	"github.com/go-macaron/macaron"
)


var(
	lock LockRedis
)

func main() {
	client:=ExampleNewClient()
	defer client.Close()
	lock=LockRedis{client}
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
			return fmt.Errorf("zero")
		}
		fmt.Println("current:",v)
		//return tx.Decr(k).Err()
		//status :=tx.Set(k,v-1,-1)
		//return status.Err()
		status:=tx.GetSet(k,v-1)
		fmt.Println(status)
		old,err2:=status.Int64()
		if err2!=nil{
			return err2
		}

		if old == v{
			return nil
		}
		return fmt.Errorf("currence!")
	},k)
	fmt.Println("watchErr:",watchErr)
	return nil== watchErr
}