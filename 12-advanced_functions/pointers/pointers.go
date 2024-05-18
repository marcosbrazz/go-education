package main

import "fmt"

// pass param by value (copy of original variable)
func invert(number int) int {
	return number * -1
}

// pass param by reference
func invertByReference(number *int) {
	*number = *number * -1
}

func main() {
	number := 2
	inverted := invert(number)
	fmt.Println(fmt.Sprintf("By value: number %d inverted %d", number, inverted))

	invertByReference(&number)

	fmt.Println(fmt.Sprintf("By reference: number %d", number))
}
