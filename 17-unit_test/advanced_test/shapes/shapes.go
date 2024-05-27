package shapes

import (
	"math"
)

// Interface (only functions can be defined)
type Shape interface {
	area() float64
}

type Circle struct {
	Ray float64
}

// Implements interface method
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Ray, 2)
}

type Rectangle struct {
	X float64
	Y float64
}

// Implements interface method
func (r Rectangle) Area() float64 {
	return r.X * r.Y
}
