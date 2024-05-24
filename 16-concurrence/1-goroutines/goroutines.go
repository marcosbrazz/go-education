package main

import (
	"fmt"
	"time"
)

// Run workloads in parallel
func main() {
	go workload("work 1") // Don't wait this call to return to continue next lines.
	workload("work 2")
}

func workload(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
