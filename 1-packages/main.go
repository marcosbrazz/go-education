package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Writing from main")
	auxiliar.Escrever()
	validation := checkmail.ValidateFormat("mumia")
	fmt.Println(validation)

}
