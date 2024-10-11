package tests

import (
	"fmt"
	"math/rand"
	"one-million-checkboxes/controllers/checkboxes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCurrentBitSet(t *testing.T) {
	bitset := checkboxes.GetCurrentBitSet()
	fmt.Println(bitset)
	assert.NotNil(t, bitset)
}

func TestUpdateBitSet(t *testing.T) {
	bit := 1
	value := true
	checkboxes.UpdateBitSet(bit, value)
	bitset := checkboxes.GetCurrentBitSet()
	fmt.Println(bitset)
	assert.NotNil(t, bitset)
}

func TestUpdateBunkBitSet(t *testing.T) {
	randomPosition := rand.Int() % 1000000
	value := true
	checkboxes.UpdateBitSet(randomPosition, value)
	bitset := checkboxes.GetCurrentBitSet()
	assert.NotNil(t, bitset)
}
