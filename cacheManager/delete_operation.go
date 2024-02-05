package cacheManager

import (
	"bufio"
	"os"
)

// Remove the cached data by uuid from slice of strings
func removeCacheByUuidIfExist(lines []string, uuid string) []string {
	for i := 0; i < len(lines); i++ {
		if lines[i] == "["+uuid+"]" {

			lines[i] = ""
			lines[i+1] = ""
			lines[i+2] = ""
			break
		}
	}
	return lines
}

// Describes a slice of strings of file content
func fileContentOf(f *os.File) []string {
	scaner := bufio.NewScanner(f)
	var lineas []string
	for scaner.Scan() {
		lineas = append(lineas, scaner.Text())
	}

	return lineas
}

// Delete the cached data by uuid
func Uncache(uuid string) error {
	f := openOrCreateCacheFile()
	defer f.Close()
	lineas := fileContentOf(f)
	lineas = removeCacheByUuidIfExist(lineas, uuid)

	f.Seek(0, 0)
	f.Truncate(0)
	for _, linea := range lineas {
		if linea != "" {
			f.WriteString(linea + "\n")
		}
	}

	return nil
}
