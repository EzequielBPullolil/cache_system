package cacheManager

import (
	"fmt"
	"strings"
)

func handleCache(message string) {
	message = message[6:]
	uuid := message[0:strings.IndexRune(message, ' ')]
	data := []byte(message[strings.IndexRune(message, ' ')+1:])
	fmt.Println("Uuid " + uuid)
	fmt.Println("data" + string(data))

	WriteCache(uuid, data)
}

func WriteCache(uuid string, data []byte) {

}
