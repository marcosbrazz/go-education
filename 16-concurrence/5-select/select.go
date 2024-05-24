package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			channel1 <- "channel 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "channel 2"
		}
	}()

	// This way takes the performance of the slowest channel
	// for {
	// 	msg1 := <-channel1
	// 	fmt.Println(msg1)

	// 	msg2 := <-channel2
	// 	fmt.Println(msg2)
	// }

	// This way doesn't block any channel receival
	for {
		select { // check which channel has a message ready to be read.
		case msg1 := <-channel1: // if message arrive in channel1, instantly execute this block
			fmt.Println(msg1)
		case msg2 := <-channel2: // if message arrive in channel2, instantly execute this block
			fmt.Println(msg2)
		}

	}
}
