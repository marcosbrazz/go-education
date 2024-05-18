package main

import "fmt"

func divide(n1, n2 float32) float32 {
	defer recover_division_from_zero()
	if n2 == 0 {
		panic("CANNOT DIVIDE TO ZERO") // If there is no recover, defer funcs are called and program is shutdown.
	}
	fmt.Println("Divide", n1, n2)
	return n1 / n2
}

func recover_division_from_zero() {
	if r := recover(); r != nil {
		fmt.Println("Recovered")
	}
}

func main() {
	result := divide(3, 2)
	fmt.Println("End", result)
}
