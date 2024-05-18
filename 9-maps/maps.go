package main

import "fmt"

func main() {

	user := map[string]string{
		"name":  "Marcos",
		"skill": "Go developer",
	}

	fmt.Println(user)

	// nested map
	user2 := map[string]map[string]string{
		"name": {
			"first": "Marcos",
			"last":  "Braz",
		},
	}

	fmt.Println(user2)

}
