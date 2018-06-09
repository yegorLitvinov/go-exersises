package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(500 * time.Millisecond)
	go func() {
		for s := range t.C {
			fmt.Println(s)
		}
	}()
	time.Sleep(5 * time.Second)
	t.Stop()
}
