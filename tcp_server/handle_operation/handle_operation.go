package handleoperation

import (
	"strings"

	"github.com/EzequielBPullolil/cache_system/cacheManager"
)

/*
handles which operation to call based on the client request
and returns a response if necessary


*/

func HandleOperation(d []byte) {
	data := strings.Split(string(d), " ")
	if len(data) >= 2 {
		operation := data[0]
		uuid := data[1]
		switch strings.ToLower(operation) {
		case "cache":
			var data_ string
			for _, e := range data[2:] {
				data_ += string(e)
			}
			cacheManager.WriteCache(uuid, data_)
		case "uncache":
			cacheManager.Uncache(uuid)
		case "fetch":
		}
	}
}
