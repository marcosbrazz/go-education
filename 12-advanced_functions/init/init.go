package main

import "fmt"

// Can have one per package
func main() {
	fmt.Println("main")
}

// Executes before main
// Can have one per file
func init() {
	fmt.Println("init")
}
