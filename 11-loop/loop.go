package main

import (
	"fmt"
	"time"
)

func main() {

	i := 0

	for i < 10 {
		i++
		fmt.Println("Increment i", i)
		time.Sleep(time.Second)
	}

	for j := 0; j < 10; j++ {
		fmt.Println("Increment j", j)
		time.Sleep(time.Second)
	}
}
