package ratelimit

import (
	"fmt"
	"testing"
	"time"
)

var (
	rate = &Rate{
		Rate:  2,
		LType: Per30secLimit,
	}
)

func init() {
	rate.Init()
}

func TestA(t *testing.T) {
	go func() {
		for {
			ok, err := rate.Req()
			fmt.Println(ok, err)
			time.Sleep(1e8)
		}
	}()
	time.Sleep(10e9)
}
