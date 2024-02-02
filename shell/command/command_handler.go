package command

import (
	"net"
	"strings"
)

func CommandHandler(s string, conn net.Conn) {
	switch s[:strings.IndexRune(s, ' ')] {
	case "Cache":

	case "Uncache":

	case "Fetch":

	}

}
