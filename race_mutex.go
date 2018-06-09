package main

import (
	"fmt"
	"sync"
	"time"
)

func co(counter map[int]int, i int, mu *sync.Mutex) {
	for j := 0; j < 2; j++ {
		time.Sleep(time.Second)
		mu.Lock()
		counter[i*10+j]++
		mu.Unlock()
	}
}

func main2() {
	mu := &sync.Mutex{}
	counter := make(map[int]int)
	for i := 0; i < 2; i++ {
		go co(counter, i, mu)
	}

	mu.Lock()
	fmt.Println(counter)
	mu.Unlock()
}
