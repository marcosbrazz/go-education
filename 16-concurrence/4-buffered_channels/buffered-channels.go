package main

import "fmt"

func main() {
	channel := make(chan string, 2) // channel with capacity of 2

	// It will block only after buffer capacity is reached.
	channel <- "first"
	channel <- "second"
	// channel <- "third"  !! it would block here !!
	msg1 := <-channel
	msg2 := <-channel

	fmt.Println(msg1)
	fmt.Println(msg2)
}
