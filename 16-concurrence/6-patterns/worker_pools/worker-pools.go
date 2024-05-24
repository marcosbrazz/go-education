package main

import "fmt"

func fibonacci(index int) int {
	if index <= 1 {
		return index
	}

	return fibonacci(index-2) + fibonacci(index-1)
}

// task channel only output data. results channel only receive data
func worker(tasks <-chan int, results chan<- int) {
	for number := range tasks {
		results <- fibonacci(number)
	}
}

func main() {
	tasks, results := make(chan int, 45), make(chan int, 45)

	// Workers consume task channel in parallel.
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)

	for i := 0; i < 45; i++ {
		tasks <- i
	}

	close(tasks)

	for i := 0; i < 45; i++ {
		fmt.Println(<-results)
	}

}
