package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/EzequielBPullolil/cache_system/cacheManager"
)

func handleCache(message string) (string, []byte) {
	message = message[6:]
	uuid := message[0:strings.IndexRune(message, ' ')]
	data := []byte(message[strings.IndexRune(message, ' ')+1:])
	fmt.Println("Uuid " + uuid)
	fmt.Println("data" + string(data))
	return uuid, data
}
func handleFetch(message string) string {
	message = message[5:]
	return message
}
func handleUncache(message string) string {
	message = message[7:]

	return message
}
func handleOperation(data []byte) {
	message := string(data)
	if strings.Contains(message, "Cache") {
		uuid, data := handleCache(message)
		cacheManager.WriteCache(uuid, data)
	} else if strings.Contains(message, "Uncache") {
		uuid := handleUncache(message)
		cacheManager.Uncache(uuid)
	} else if strings.Contains(message, "Fetch") {
		uuid := handleFetch(message)
		cacheManager.Fetch(uuid)
	} else {
		fmt.Println("No operation")
	}
}

func main() {
	listen, err := net.Listen("tcp", "localhost:29033")

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer listen.Close()

	for {
		data := make([]byte, 1024)
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		clientIP, port, _ := net.SplitHostPort(conn.RemoteAddr().String())
		fmt.Println("Connection created " + clientIP + ":" + port)
		conn.Read(data)
		go handleOperation(data)
		conn.Close()
	}
}
