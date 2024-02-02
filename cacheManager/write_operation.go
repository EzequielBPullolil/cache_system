package cacheManager

import (
	"os"
)

// Write cache file associating uuid with data
func WriteCache(uuid string, data []byte) error {
	//cache_data := CacheData(uuid, data, "string")
	f, _ := os.ReadFile("cache.txt")

	f = append(f, []byte(uuid)...)

	os.WriteFile("cache.txt", f, os.ModePerm)
	return nil
}
