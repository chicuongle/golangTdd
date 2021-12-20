package functions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeEvenGenerator() func() uint {
	init := uint(0)
	return func() uint {
		ret := init
		init += 2
		return ret
	}
}

func TestEvenGenerator(t *testing.T) {
	t.Run("function returned by makeEvenGenerator is combined with variable init. the init varialbe is persisted between calls", func(t *testing.T) {
		nextEven := makeEvenGenerator()
		assert.Equal(t, uint(0), nextEven())
		assert.Equal(t, uint(2), nextEven())
		assert.Equal(t, uint(4), nextEven())
	})
}
