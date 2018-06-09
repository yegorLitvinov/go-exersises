package main

import "fmt"

func main1() {
	messages := make(chan string)
	go func() {
		messages <- "Hello"
	}()
	fmt.Println(<-messages)
}
