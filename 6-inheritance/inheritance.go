package main

import "fmt"

type person struct {
	name   string
	age    uint8
	height uint8
}
type student struct {
	person // inherits attributes from the person struct
	course string
}

func main() {
	p1 := person{"Marcos", 40, 181}
	fmt.Println(p1)
	p2 := student{p1, "computer science"}
	fmt.Println(p2)
	fmt.Println(p2.name)
}
