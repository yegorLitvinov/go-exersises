package main

import (
	"fmt"
	"sync/atomic"
)

var number int32

func inc() {
	atomic.AddInt32(&number, 1)
}

func main4() {
	for i := 0; i < 100; i++ {
		go inc()
	}
	// time.Sleep(time.Second)
	fmt.Println(atomic.LoadInt32(&number))
}
