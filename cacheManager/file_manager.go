package cacheManager

import (
	"os"
)

type jsonFile struct {
	elements map[string]string
}

func openOrCreateCacheFile() *os.File {
	file, _ := os.OpenFile(".cache_file", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	return file
}

func appendCache(f *os.File, d cacheData) error {
	f.Write([]byte(d.format()))
	f.Close()
	return nil
}
