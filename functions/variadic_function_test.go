package functions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
number of input parameters is not fixed, this number can be variable and depends on function consumer
*/
func sum(args ...int) int {
	total := 0
	for _, val := range args {
		total += val
	}
	return total
}

func TestVariadicFunc(t *testing.T) {
	t.Run("call variadic function with any number of inputs", func(t *testing.T) {
		assert.Equal(t, 28, sum(1, 2, 3, 4, 5, 6, 7))
	})
	t.Run("call variadic function with a slice as input", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5, 6, 7}
		assert.Equal(t, 28, sum(slice...))
	})
}
