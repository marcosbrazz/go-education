package main

import "fmt"

func query_database() {
	defer fmt.Println("close connection") // the last sentence to execute b4 return or end function
	fmt.Println("open connection")
	fmt.Println("execute query")
	fmt.Println("return result or raise error")
}

func main() {
	query_database()
}
