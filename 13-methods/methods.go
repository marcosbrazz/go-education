package main

import "fmt"

type user struct {
	name string
	age  uint8
}

// getter method
func (u user) getName() string {
	return u.name
}

// setter method
func (u *user) setName(name string) {
	u.name = name
}

func main() {
	myUser := user{"Marcos", 40}
	fmt.Println(myUser)

	myUser.setName("Braz")
	fmt.Println(myUser)
	fmt.Println(myUser.getName())

}
