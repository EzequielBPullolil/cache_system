package cacheManager

import (
	"fmt"
	"strings"
)

func HandleOperation(data []byte) {
	message := string(data)
	if strings.Contains(message, "Cache") {
		handleCache(message)
	} else if strings.Contains(message, "Uncache") {
		handleUncache(message)
	} else if strings.Contains(message, "Fetch") {
		handleFetch(message)
	} else {
		fmt.Println("No operation")
	}
}
