package cacheManager

import "fmt"

type cacheData struct {
	Uuid      string `json:"uuid"`
	Data      string `json:"data"`
	Data_type string `json:"data_type"`
}

func CacheData(uuid string, data string, data_type string) cacheData {
	return cacheData{
		Uuid:      uuid,
		Data:      data,
		Data_type: data_type,
	}
}

// / Describe an CacheData formated to string
func (d cacheData) format() string {
	return fmt.Sprintf("[%s]\ndata=%s\nData_type=%s\n", d.Uuid, d.Data, d.Data_type)
}
