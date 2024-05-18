package main

import "fmt"

type user struct {
	name    string
	age     uint8
	address address
}
type address struct {
	street string
	number uint8
}

func main() {
	var u user
	u.age = 40
	u.name = "Marcos"
	u.address.street = "Mummy street"
	var addr address
	addr.street = "ABC street"
	addr.number = 12
	fmt.Println(u)
	fmt.Println(user{"São braz", 40, addr})
	fmt.Println(user{age: 40})
}
