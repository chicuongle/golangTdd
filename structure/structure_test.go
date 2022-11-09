package structure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicStructure(t *testing.T) {
	t.Run("basic characters", func(t *testing.T) {
		got := Rectangle{10.0, 10.0}
		assert.Equal(t, 40.0, got.Area())
	})
	t.Run("data driven tests using interface", func(t *testing.T) {
		areaTests := []struct {
			shape Shape
			want  float64
		}{
			{Rectangle{12, 6}, 72.0},
			{Circle{10}, 314.1592653589793},
			{Triangle{12, 6}, 36.0},
		}
	
		for _, tt := range areaTests {
			assert.Equal(t,tt.want,tt.shape.Area())
		}
	})
}
