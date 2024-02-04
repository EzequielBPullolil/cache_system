package cacheManager

import (
	"os"
	"strings"
)

// Iterate the file finding an uuid equal to 'uuid'
func FetchCache(uuid string) (string, error) {
	d, _ := os.ReadFile(os.Getenv("cache_file"))
	data := ""
	file_data := strings.Split(string(d), "\n")
	for i := 0; i < len(file_data); i++ {
		if file_data[i] == "["+uuid+"]" {
			data = file_data[i+1]
		}
	}
	return data, nil

}
