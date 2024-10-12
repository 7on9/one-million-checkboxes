package bit_sets

import (
	"context"
	"log"
	"one-million-checkboxes/services/redis"

	"github.com/bits-and-blooms/bitset"
)

const (
	MaxBitSetLength = 64
	KeyBitSet       = "CURRENT_BITSET"
)

var (
	ctx    = context.Background()
	client = redis.GetRedisClient()
	bs     = bitset.New(MaxBitSetLength)
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
		// client.SetBit(ctx, KeyBitSet, 0, 0)
	} else {
		log.Println("Bitset found in Redis")
		// Load the bitset from Redis
		// bitfield := client.GetBit(ctx, KeyBitSet, 0).Val()
		// for i := 0; i < MaxBitSetSize; i++ {
		// 	bit := client.GetBit(ctx, KeyBitSet, int64(i)).Val()
		// 	if bit == 1 {
		// 		bs.Set(uint(i))
		// 	}
		// }

		// arrayGetBitCmd := make([][]string, 0)
		// for i := 0; i < MaxBitSetSize; i++ {
		// 	arrayGetBitCmd = append(arrayGetBitCmd, []string{"get", fmt.Sprintf("u%d", i/2), fmt.Sprintf("%d", i)})
		// }
		// bitFields := client.BitField(ctx, KeyBitSet, arrayGetBitCmd).Val()
		// for i, bitField := range bitFields {
		// 	if bitField == 1 {
		// 		bs.Set(uint(i))
		// 	}
		// }
	}
}

func GetCurrentBitSet() string {
	log.Printf("Getting the current bitset %d \n", bs.Len())
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

func WipeBitSet() {
	log.Println("Wiping the bitset")
	client.Del(ctx, KeyBitSet)
	bs.ClearAll()
}
