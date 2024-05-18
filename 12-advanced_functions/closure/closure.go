package main

import "fmt"

func closure() func() {
	text := "Inside closure"

	function := func() {
		fmt.Println(text)
	}

	return function
}

func main() {

	text := "Inside main"
	fmt.Println(text)

	my_func := closure()

	my_func()

}
