package command

import (
	"net"
	"strings"
)

func handleFetch(s string, conn net.Conn) {
	message := s[5:]
	uuid := message[strings.IndexRune(message, ' '):]

	uuid = strings.TrimSpace(uuid)

	conn.Write([]byte("Fetch " + uuid))
}

/// Cache asdasdasda-asdasda {"hola":"hola"}
