package main

import "fmt"

func main() {
	var variable int
	var pointer *int // '*' => pointer data type

	variable = 100
	pointer = &variable // & => access variable memory address

	fmt.Println(variable, pointer) // variable value, memory address
	fmt.Println(*pointer)          // variable value
}
