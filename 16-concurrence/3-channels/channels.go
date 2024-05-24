package main

import (
	"fmt"
	"time"
)

// Run workloads in parallel
func main() {
	myChannel := make(chan string) // Create a channel to exchange string
	go workload("workload 1", myChannel)

	fmt.Println("After workload 1 started")

	for {
		text, isOpen := <-myChannel // Wait until channel receive data.
		if !isOpen {                // isOpen: channel flag to indicate whether the channel is closed
			break
		}
		fmt.Println(text)
	}

	myChannel = make(chan string) // Create a channel to exchange string
	go workload("workload 2", myChannel)

	fmt.Println("After workload 2 started")

	// Another way to do the previous "for"
	// while channel is open, print the text.
	for text := range myChannel {
		fmt.Println(text)
	}

}

func workload(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text // Send text data to channel
		time.Sleep(time.Second)
	}
	close(channel) // Close the channel when no more data will be sent through it.
}
