package functions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValueReceiver(t *testing.T) {
	t.Run("parameter is copied", func(t *testing.T) {
		myIllusion := illusion{
			magicianCount: 0,
		}
		myIllusion.increaseMagicianCount()
		assert.NotEqual(t, 1, myIllusion.magicianCount, "after increased magic count, properties of struct is expected to be increased but not ...")
		assert.Equal(t, 0, myIllusion.magicianCount, "magic count still unchanged")
	})
	t.Run("a function applied to illusion", func(t *testing.T) {
		myIllusion := illusion{
			magicianCount: 0,
		}
		increaseMagicCount(myIllusion)
		assert.Equal(t, 0, myIllusion.magicianCount, "after applied the util increase function to illusion, the magic count is unchanged. "+
			"A new illusion struct is copied in function executing heap, and therefore the original illusion struct is not changed")
	})
}
