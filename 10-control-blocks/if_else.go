package main

import "fmt"

func main() {

	number := 10

	if number >= 10 {
		fmt.Println("Bigger or equal to 10")
	}

	// if init (variable scoped to if / else block only)
	if increment_of_ten := number + 2; increment_of_ten > 10 {
		fmt.Println("Bigger than 10")
	}

}
