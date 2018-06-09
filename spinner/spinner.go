package main

import (
	"fmt"
	"time"
)

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

func spinner() {
	for {
		for _, n := range "\\|/-" {
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("\r%c", n)
		}
	}
}

func main() {
	go spinner()
	fmt.Printf("\rResult is %d\n", fib(43))
}
