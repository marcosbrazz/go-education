package main

import "fmt"

func main() {

	var array1 [5]string
	array1[0] = "Pos 1"
	fmt.Println(array1)

	array2 := [3]string{"Pos 1", "Pos 2", "Pos 3"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5} // Size dynamically defined by content
	fmt.Println(array3)

	slice := []int{6, 7, 8, 9} // slice works as pointer to an array. Resizable.
	fmt.Println(slice)

	slice = append(slice, 10) // Recreate slice with new item
	fmt.Println(slice)

	slice2 := array3[1:3] // Create a slice from array position 1 to 2
	fmt.Println(slice2)

	array3[1] = 99 // Changed value in array reflect in slice, because slices are pointers.
	fmt.Println(slice2)

	// Internal Arrays

	println("--- Internal Arrays ----")

	slice3 := make([]float32, 10, 11) // Type, size, capacity
	fmt.Println(slice3)
	fmt.Println(len(slice3), cap(slice3))

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6) // Created a new array with double of capacity
	fmt.Println(len(slice3), cap(slice3))

	slice4 := make([]float32, 5) // Omit capacity, capacity = size
	fmt.Println(len(slice4), cap(slice4))

}
