package bit_sets

import (
	"context"
	"log"
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

func InitBitSetsServices() {
	// Initialize the bitset
	// Retrieve the bitset from Redis
	// If the bitset is not found, create a new one
	// If the bitset is found, load it into the bitset
	isBitSetExist := client.Exists(ctx, KeyBitSet).Val()

	if isBitSetExist == 0 {
		log.Println("Bitset not found in Redis")
		// Create a new bitset
		client.SetBit(ctx, KeyBitSet, 0, 0)
	} else {
		log.Println("Bitset found in Redis")
		// Load the bitset from Redis
		bitfield := client.BitField(ctx, KeyBitSet, "GET", "u1", 0).Val()
		for i, bit := range bitfield {
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
	log.Printf("Updating bitset at position %d with value %t \n", bit, value)
	if value {
		bs.Set(uint(bit))
		client.SetBit(ctx, KeyBitSet, int64(bit), 1)
	} else {
		bs.Clear(uint(bit))
		client.SetBit(ctx, KeyBitSet, int64(bit), 0)
	}
}
