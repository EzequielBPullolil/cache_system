package cacheManager

import (
	"encoding/json"
	"os"
)

type jsonFile struct {
	elements map[string]string
}

func openOrCreateCacheFile() *os.File {
	file, _ := os.OpenFile(os.Getenv("cache_file"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	return file
}

func appendCache(f *os.File, d cacheData) error {
	encodedData, _ := json.MarshalIndent(d, "", "    ")
	f.Write(encodedData)
	defer f.Close()
	return nil
}
