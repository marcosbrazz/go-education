package main

import (
	"errors"
	"fmt"
)

func main() {
	// int8, int16, int32, int64
	var n1 int8 = 123
	// infers bit width based on machine architecture
	var n3 int = 789
	n2 := 456
	var n4 uint8 = 100    // unsigned int
	var n5 byte = 101     // alias for uint8
	var n6 rune = 1000000 // alias for int
	// float32, float64
	var n7 float32 = 123.456
	n8 := 33.56 // infers bit width based on machine architecture
	n8 = 3
	var b2 bool = true
	var err2 error = errors.New("Err msg")
	char := 'A' // converts to ASCII table number
	// Initial values
	var empty string  // empty string
	var zero int      // 0
	var zerof float32 // 0
	var b1 bool       // false
	var err1 error    // nil
	fmt.Println(n1, n2, n3, n4, n5, n6, n7, n8, char, b2, err2)
	fmt.Println(empty)
	fmt.Println(zero, zerof, b1, err1)
}
