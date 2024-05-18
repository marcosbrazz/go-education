package main

import "fmt"

// Vars declaration
func main() {
	var var1 string = "var1"
	var2 := "var2"
	var (
		var3 string = "var3"
		var4 string = "var4"
	)
	var5, var6 := "var5", "var6"

	var5, var6 = var6, var5

	const var7 string = "var7"

	fmt.Println(var1+var2, var3, var4, var5, var6, var7)
}
