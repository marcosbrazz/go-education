package main

import "fmt"

// Name the variables will be returned
func math(n1, n2 int) (sum int, diff int) {
	sum = n1 + n2
	diff = n1 - n2
	return
}

func main() {
	s, d := math(5, 2)
	fmt.Println(s, d)
}
