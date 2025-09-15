package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hello, world!"
	}()

	time.Sleep(1 * time.Second)

	select {
	case msg := <-ch:
		fmt.Println(msg)
	default:
		fmt.Println("No message received")
	}
}
