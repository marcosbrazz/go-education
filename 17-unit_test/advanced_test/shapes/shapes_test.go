package shapes

import (
	"math"
	"testing"
)

func TestShapes(t *testing.T) {

	t.Run("rectangle", func(t *testing.T) {
		ret := Rectangle{10.0, 5.0}
		actual := ret.Area()
		expected := 50.0
		if actual != expected {
			t.Error("Test Failed")
		}

	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{1}
		expected := math.Pi
		actual := circle.Area()

		if actual != expected {
			t.Error("Test Failed")
		}
	})
}
