package main

import (
	"fmt"
	"time"
)

func main() {
	channel := generateText("hello")
	for i := 0; i < 5; i++ {
		fmt.Println(<-channel)
	}
}

func generateText(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("received value %s", text)
			time.Sleep(time.Second)
		}
	}()

	return channel

}
