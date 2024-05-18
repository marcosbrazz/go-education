package main

import "fmt"

func fibonacci(index uint) uint {
	if index <= 1 {
		return index
	}

	return fibonacci(index-2) + fibonacci(index-1)
}

func main() {

	// 1 1 2 3 5 8 13...
	index := uint(13)
	fmt.Println(fibonacci(index))
}
