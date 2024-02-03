package command

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func CommandHandler(s string, conn net.Conn) {
	if strings.Contains(s, "Cache") {
		handleCache(s, conn)
	} else if strings.Contains(s, "Fetch") {
		handleFetch(s, conn)
	} else if strings.Contains(s, "Uncache") {

	} else if strings.Contains(s, "Exit") || strings.Contains(s, "exit") {
		conn.Close()
		fmt.Println("Exit")
		os.Exit(1)
	}
}
