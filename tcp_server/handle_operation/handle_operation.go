package handleoperation

import (
	"strings"

	"github.com/EzequielBPullolil/cache_system/cacheManager"
)

/*
handles which operation to call based on the client request
and returns a response if necessary
*/

func HandleOperation(d []byte, response *string) {
	data := strings.Split(string(d), " ")
	if len(data) >= 2 {
		operation := data[0]
		uuid := data[1]
		switch strings.ToLower(operation) {
		case "cache":
			cacheManager.WriteCache(uuid, data[2])
		case "uncache":
			cacheManager.Uncache(uuid)
		case "fetch":
			cacheManager.FetchCache(uuid)
		}
	} else {
		*response = "invalid operation"
	}
}
