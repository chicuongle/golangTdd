package pointer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWallet(t *testing.T) {
	wallet:= Wallet{}
	wallet.Deposit(10.0)
	assert.Equal(t,10.0,wallet.Balance())
}
