package command

import (
	"fmt"
	"net"
	"strings"
)

func handleCache(s string, conn net.Conn) {
	message := s[6:]
	uuid := message[0:strings.IndexRune(message, ' ')]
	data := message[strings.IndexRune(message, ' ')+1:]

	uuid = strings.TrimSpace(uuid)
	data = strings.TrimSpace(data)

	fmt.Print("Cached: ")
	fmt.Println(uuid, data)

	conn.Write([]byte("Cache " + uuid + " " + data))
}

/// Cache asdasdasda-asdasda {"hola":"hola"}
