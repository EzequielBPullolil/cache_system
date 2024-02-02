package cacheManager

// Write cache file associating uuid with data
func WriteCache(uuid string, data string) error {
	cd := CacheData(uuid, data, "string")
	f := openOrCreateCacheFile()
	return appendCache(f, cd)
}
