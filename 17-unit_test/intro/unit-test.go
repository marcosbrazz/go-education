package main

import (
	"fmt"
	"intro/address"
)

func main() {
	fmt.Println(address.AddressType("Main St"))
	fmt.Println(address.AddressType("Kirkman Road"))
	fmt.Println(address.AddressType("Kirkman Zt"))
}
