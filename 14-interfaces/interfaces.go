package main

import (
	"fmt"
	"math"
)

// Interface (only functions can be defined)
type shape interface {
	area() float64
}

type circle struct {
	ray float64
}

// Implements interface method
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.ray, 2)
}

type retangle struct {
	x float64
	y float64
}

// Implements interface method
func (r retangle) area() float64 {
	return r.x * r.y
}

// Function that receives an interface type
func printArea(s shape) {
	fmt.Printf("Area %0.2f\n", s.area())
}

func main() {
	r := retangle{3, 5}
	printArea(r)

	c := circle{66}
	printArea(c)
}
