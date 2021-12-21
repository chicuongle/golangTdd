package datastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Rectangle struct {
	heigth float64
	width  float64
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

func (r *Rectangle) Area() float64 {
	return r.heigth * r.width
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.heigth + r.width)
}
