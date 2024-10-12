package tests

import (
	"fmt"
	"math/rand"
	bitSets "one-million-checkboxes/services/bit_sets"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCurrentBitSet(t *testing.T) {
	bitset := bitSets.GetCurrentBitSet()
	fmt.Println(bitset)
	assert.NotNil(t, bitset)
}

func TestUpdateBitSet(t *testing.T) {
	bit := 1
	value := true
	bitSets.UpdateBitSet(bit, value)
	bitset := bitSets.GetCurrentBitSet()
	fmt.Println(bitset)
	assert.NotNil(t, bitset)
}

func TestUpdateBunkBitSet(t *testing.T) {
	randomPosition := rand.Int() % 1000000
	value := true
	bitSets.UpdateBitSet(randomPosition, value)
	bitset := bitSets.GetCurrentBitSet()
	assert.NotNil(t, bitset)
}
