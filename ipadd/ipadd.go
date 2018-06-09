package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: ipadd 127.0.0.1\n")
		os.Exit(1)
	}
	ip := net.ParseIP(os.Args[1])
	if ip == nil {
		fmt.Println("Invalid Address")
	} else {
		fmt.Printf("Address is %s\n", ip)
	}
	os.Exit(0)
}
