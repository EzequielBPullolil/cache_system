package command

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func CommandHandler(s string, conn net.Conn) {
	if strings.Contains(s, "Exit") || strings.Contains(s, "exit") {
		conn.Close()
		fmt.Println("Exit")
		os.Exit(1)
	}

	conn.Write([]byte(s))
}
