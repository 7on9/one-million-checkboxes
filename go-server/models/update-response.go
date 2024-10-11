package models

import (
	"github.com/bits-and-blooms/bitset"
)

type UpdateResponse struct {
	BitSet bitset.BitSet `json:"bitSet"`
}
