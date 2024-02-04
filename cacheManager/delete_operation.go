package cacheManager

import (
	"bufio"
)

func Uncache(uuid string) error {
	f := openOrCreateCacheFile()
	defer f.Close()
	scaner := bufio.NewScanner(f)
	var lineas []string
	for scaner.Scan() {
		lineas = append(lineas, scaner.Text())
	}
	for i := 0; i < len(lineas); i++ {
		if lineas[i] == "["+uuid+"]" {

			lineas[i] = ""
			lineas[i+1] = ""
			lineas[i+2] = ""
			break
		}
	}

	f.Seek(0, 0)
	f.Truncate(0)
	for _, linea := range lineas {
		if linea != "" {
			f.WriteString(linea + "\n")
		}
	}

	return nil
}
