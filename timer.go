package main

import (
	"fmt"
	"time"
)

func main6() {
	t1 := time.NewTimer(time.Second)
	t2 := time.NewTimer(2 * time.Second)

	<-t1.C
	fmt.Println("t1 expired")

	go func() {
		<-t2.C
		fmt.Println("t2 expired")
	}()

	if t2.Stop() {
		fmt.Println("t2 stopped")
	}
}
