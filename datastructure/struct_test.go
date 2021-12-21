package datastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerimeter(t *testing.T) {
	t.Run("check perimeter", func(t *testing.T) {
		assert.Equal(t, 40.0, Perimeter(10.0, 10.0))
	})
	t.Run("check area calcualtion", func(t *testing.T) {
		assert.Equal(t, 72.0, Area(12.0, 6))
	})
}

func Area(heigth float64, width float64) float64 {
	return heigth * width
}

func Perimeter(heigth, width float64) float64 {
	return 2 * (heigth + width)
}
