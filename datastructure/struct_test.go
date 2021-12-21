package datastructure

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type shape interface {
	Area() float64
	Perimeter() float64
}

func getArea(s shape) float64 {
	return s.Area()
}

type Rectangle struct {
	heigth float64
	width  float64
}

func (r *Rectangle) Area() float64 {
	return r.heigth * r.width
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.heigth + r.width)
}

type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func TestPerimeter(t *testing.T) {
	t.Run("check perimeter", func(t *testing.T) {
		rectangle := &Rectangle{10.0, 10.0}
		assert.Equal(t, 40.0, rectangle.Perimeter())
	})
	t.Run("check area calcualtion", func(t *testing.T) {
		rectangle := &Rectangle{12.0, 6}
		assert.Equal(t, 72.0, rectangle.Area())
	})
}
func TestShape(t *testing.T) {
	t.Run("using for rectangle", func(t *testing.T) {
		rectangle := &Rectangle{10.0, 10.0}
		assert.Equal(t, 100.0, getArea(rectangle))
	})
	t.Run("using for a circle", func(t *testing.T) {
		rectangle := &Circle{10.0}
		assert.InEpsilon(t, 314.15, getArea(rectangle), float64(0.001))
	})
}
