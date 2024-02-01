package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

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

func handleOperation(data []byte) {
	message := string(data)
	if strings.Contains(message, "Cache") {
		handleCache(message)
	} else if strings.Contains(message, "Uncache") {
		handleUncache(message)
	} else if strings.Contains(message, "Fetch") {
		handleFetch(message)
	} else {
		fmt.Println("No operation")
	}
}
func handleFetch(data string) {

}
func handleCache(message string) {
	message = message[6:]
	uuid := message[0:strings.IndexRune(message, ' ')]
	data := []byte(message[strings.IndexRune(message, ' ')+1:])
	fmt.Println("Uuid " + uuid)
	fmt.Println("data" + string(data))

	writeCache(uuid, data)

}
func handleUncache(data string) {

}

func writeCache(uuid string, data []byte) {

	f, _ := os.ReadFile("cache.txt")

	if f == nil {
		file, _ := os.Create("cache.txt")

		file.Write([]byte(uuid))
		file.Sync()
	}
}
