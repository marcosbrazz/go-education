package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cat struct {
	Name  string `json:"cat_name"`
	Color string `json:"color_cat"`
	Age   uint8  `json:"age"`
}

func main() {
	d := cat{"Toby", "gray", 8}

	dJson, error := json.Marshal(d)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(dJson)
	fmt.Println(bytes.NewBuffer(dJson))

	otherCat := map[string]string{
		"name":  "pipoca",
		"color": "gold",
	}

	otherJson, error := json.Marshal(otherCat)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(bytes.NewBuffer(otherJson))

}
