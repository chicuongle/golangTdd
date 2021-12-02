package functions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointerReceiver(t *testing.T) {
	t.Run("call function like a behaviour of object/struct", func(t *testing.T) {
		myIllusion := illusion{}
		myIllusion.increaseMagicianCountRef()
		assert.Equal(t, 1, myIllusion.magicianCount, "attributes of illusion is changed after increased function with pointer receiver called")
	})
	t.Run("call function like a behaviour of object/struct - other style to declared struct at runtime", func(t *testing.T) {
		myIllusion := &illusion{}
		myIllusion.increaseMagicianCountRef()
		assert.Equal(t, 1, myIllusion.magicianCount, "attributes of illusion is changed after increased function with pointer receiver called")
	})

	t.Run("call function like util function", func(t *testing.T) {
		myIllusion := illusion{}
		increaseMagicianCountRef(&myIllusion)
		assert.Equal(t, 1, myIllusion.magicianCount, "attributes of illusion is changed after increased function with pointer receiver called")
	})
}
