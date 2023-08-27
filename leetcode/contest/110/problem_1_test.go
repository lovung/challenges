package contest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_accountBalanceAfterPurchase(t *testing.T) {
	assert.Equal(t, 90, accountBalanceAfterPurchase(9))
	assert.Equal(t, 80, accountBalanceAfterPurchase(15))
	assert.Equal(t, 90, accountBalanceAfterPurchase(14))
}
