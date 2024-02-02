package cacheManager

import "strings"

type cacheData struct {
	Uuid      string `json:"uuid"`
	Data      string `json:"data"`
	Data_type string `json:"data_type"`
}

func CacheData(uuid string, data string, data_type string) cacheData {
	return cacheData{
		Uuid:      strings.TrimSpace(uuid),
		Data:      strings.TrimSpace(data),
		Data_type: data_type,
	}
}
func semiconsData(d string) string {
	return string('"') + d + string('"')
}
