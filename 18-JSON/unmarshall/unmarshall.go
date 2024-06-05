package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cat struct {
	Name  string `json:"-"` // ignored when unmarshalled
	Color string `json:"color_cat"`
	Age   uint8  `json:"age"`
}

func main() {
	catJSON := `{"cat_name":"Toby","color_cat":"gray","age":8}`
	otherCat := `{"color":"gold","name":"pipoca"}`

	var c cat

	if error := json.Unmarshal([]byte(catJSON), &c); error != nil {
		log.Fatal(error)
	}

	fmt.Println(c)

	// Unmarshal to map
	cmap := make(map[string]string)
	if error := json.Unmarshal([]byte(otherCat), &cmap); error != nil {
		log.Fatal(error)
	}

	fmt.Println(cmap)
}
