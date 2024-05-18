package main

import "fmt"

func main() {

	// anonymous function
	sum := func(n1 int, n2 int) string {
		return fmt.Sprintf("Sum: %d", n1+n2)
	}

	result := sum(5, 2)

	fmt.Println(result)

}
