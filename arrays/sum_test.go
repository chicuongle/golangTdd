package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 6, 7}
		got := Sum(numbers)
		want := 23
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	assert.Equal(t, []int{3, 9}, SumAll([]int{1, 2}, []int{0, 9}), "summarize each array and return 2 sum values")
	assert.Equal(t, []int{3, 9, 5}, SumAll([]int{1, 2}, []int{0, 9}, []int{1, 2, 2}), "number of parameters are variable")
}
