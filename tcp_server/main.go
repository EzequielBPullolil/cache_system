package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/EzequielBPullolil/cache_system/cacheManager"
)

func handleCache(message string) (string, string) {
	message = message[6:]
	uuid := message[0:strings.IndexRune(message, ' ')]
	data := message[strings.IndexRune(message, ' ')+1:]
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

	switch strings.ToLower(message[:strings.IndexRune(message, ' ')]) {
	case "cache":
		if err := cacheManager.WriteCache(handleCache(message)); err != nil {
			fmt.Println("error")
		}
	case "uncache":
		if err := cacheManager.Uncache(handleUncache(message)); err != nil {
			fmt.Println("error")
		}
	case "fetch":
		if err := cacheManager.Fetch(handleFetch(message)); err != nil {
			fmt.Println("error")
		}
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
		n, _ := conn.Read(data)

		if n > 0 {
			go handleOperation(data)
		}
		conn.Close()
	}
}
