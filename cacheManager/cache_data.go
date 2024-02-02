package cacheManager

type cacheData struct {
	uuid      string
	data      []byte
	data_type string
}

func CacheData(uuid string, data []byte, data_type string) cacheData {
	return cacheData{}
}
