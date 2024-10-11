package checkboxes

import (
	"context"
	"one-million-checkboxes/services/redis"

	"github.com/bits-and-blooms/bitset"
)

const (
	// MaxBitSetSize is the maximum size of the bitset
	MaxBitSetSize = 1000000
	KeyBitSet     = "CURRENT_BITSET"
)

var (
	ctx    = context.Background()
	client = redis.GetRedisClient()
	bs     = bitset.New(MaxBitSetSize)
)

func init() {
	// Initialize the bitset
	// Retrieve the bitset from Redis
	// If the bitset is not found, create a new one
	// If the bitset is found, load it into the bitset
	isBitSetExist := client.Exists(ctx, KeyBitSet).Val()
	if isBitSetExist == 0 {
		// Create a new bitset
		client.SetBit(ctx, KeyBitSet, 0, 0)
	} else {
		// Load the bitset from Redis
		for i := 0; i < MaxBitSetSize; i++ {
			bit := client.GetBit(ctx, KeyBitSet, int64(i)).Val()
			if bit == 1 {
				bs.Set(uint(i))
			}
		}
	}
}

func GetCurrentBitSet() string {
	return bs.DumpAsBits()
	// return bs.String()
}

func UpdateBitSet(bit int, value bool) {
	if value {
		bs.Set(uint(bit))
		client.SetBit(ctx, KeyBitSet, int64(bit), 1)
	} else {
		bs.Clear(uint(bit))
		client.SetBit(ctx, KeyBitSet, int64(bit), 0)
	}
}
