package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2))

	var f = func(txt string) string {
		return txt
	}

	println(f("some text"))

	// Multiple returns
	sum, sub := calcs(10, 15)
	println(sum, sub)

	// Ignore returned value using "_"
	sum1, _ := calcs(10, 15)
	println(sum1)

	sum2, sub2 := calcsNamed(5, 2)
	fmt.Println(sum2, sub2)

}

func sum(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// Multiple returns
func calcs(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

// NAMED RETURN
// Name the variables will be returned
func calcsNamed(n1, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}
