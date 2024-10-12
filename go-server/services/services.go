package services

import (
	bitSets "one-million-checkboxes/services/bit_sets"
	redis "one-million-checkboxes/services/redis"
)

func InitServices() {
	// InitService
	bitSets.InitBitSetsServices()
	redis.InitRedisService()
}
