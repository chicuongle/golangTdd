package functions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func simpleMultipleReturn() (int, int, int) {
	return 1, 2, 3
}

func multipleReturnTypes() (int, int, string) {
	return 1, 2, "function finish"
}

func TestReturnMultipleValue(t *testing.T) {
	t.Run("simple multiple returns", func(t *testing.T) {
		var first, second int
		first, second, _ = simpleMultipleReturn()
		assert.Equal(t, 1, first)
		assert.Equal(t, 2, second)
	})
	t.Run("return multiple values  with different types", func(t *testing.T) {
		var val string
		_, _, val = multipleReturnTypes()
		assert.Equal(t, "function finish", val)
	})
}

/*
You can name retunred values in which case they act as implicit local variable.
An empty "return" statement can then be used to return their current values.
This is known as "naked" return
*/
func inverse(v float32) (reciprocal float32) {
	if v == 0 {
		return
	}
	reciprocal = 1 / v
	return
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func TestNamedReturnValues(t *testing.T) {
	t.Run("calling a naked return function", func(t *testing.T) {
		assert.Equal(t, float32(0.5), inverse(2))
	})
	t.Run("multiple returns", func(t *testing.T) {
		resultX, resultY := split(3)
		assert.Equal(t, 1, resultX)
		assert.Equal(t, 2, resultY)
	})
}

func intAdd(a, b int) int {
	return a + b
}

func intSub(a, b int) int {
	return a - b
}

func executeFunc(a, b int, intOp func(int, int) int) int {
	fmt.Printf("intOp(%d, %d) = %d\n", a, b, intOp(a, b))
	return intOp(a, b)
}
func TestVariablesFunctionType(t *testing.T) {
	t.Run("functions can be assigned and used like variables", func(t *testing.T) {
		var intOp func(int, int) int

		intOp = intAdd
		assert.Equal(t, 2, intOp(1, 1))
		assert.Equal(t, 57, intOp(23, 34))
		intOp = intSub
		assert.Equal(t, 66, intOp(100, 34))
	})

	t.Run("functions passed as parameters", func(t *testing.T) {
		assert.Equal(t, 5, executeFunc(2, 3, intAdd))
	})

	t.Run("literals functions passed as parameters", func(t *testing.T) {
		assert.Equal(t, 6, executeFunc(2, 3, func(i1, i2 int) int {
			return i1 * i2
		}))
	})
}
