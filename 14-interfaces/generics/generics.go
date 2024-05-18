package main

import "fmt"

// This function receives any type as parameter and also return any type
func my_generic_func(arg interface{}) interface{} {
	fmt.Println(arg)
	return arg
}

func main() {
	my_generic_func(12)
	my_generic_func("text")
	my_generic_func(true)
}
