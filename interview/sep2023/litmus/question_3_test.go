package litmus

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQ3(t *testing.T) {
	var fireDragon FireDragon
	var egg ReptileEgg = fireDragon.Lay()
	var childDragon Reptile = egg.Hatch()
	assert.NotNil(t, childDragon)
	assert.Nil(t, egg.Hatch())
}
