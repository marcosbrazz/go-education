package main

import "fmt"

func sum(numbers ...int) {
	result := 0
	for _, number := range numbers {
		result += number
	}
	fmt.Println(result)
}

func sum2(title string, numbers ...int) {
	result := 0
	for _, number := range numbers {
		result += number
	}
	fmt.Println(title, result)
}

func main() {
	sum(1, 2, 55, 67, 89, 32) // Slice with n elements
	sum()                     // Empty slice
	sum2("my sum", 3, 2, 5, 1)
}
