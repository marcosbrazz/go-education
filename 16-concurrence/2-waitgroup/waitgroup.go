package main

import (
	"fmt"
	"sync"
	"time"
)

// Run workloads in parallel
func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2) // number of go routines to run.

	go func() {
		workload("go routine 1")
		waitGroup.Done()
	}()

	go func() {
		workload("go routine 2")
		waitGroup.Done()
	}()

	waitGroup.Wait() // Wait all go routines to return.

	fmt.Println("done")

}

func workload(text string) {
	for i := 0; i < 2; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
