package main

import (
	"fmt"
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
func handleOperation(data []byte, response *string) {
	var err error
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
		*response, err = cacheManager.FetchCache(handleFetch(message))
		if err != nil {
			fmt.Println("error")
		}
	}
}

// Stop app if err != nil
func checkError(err error, error_message string) {
	if err != nil {
		fmt.Println(error_message)
		os.Exit(1)
	}
}

// Create an TCP server instance using host and port --env variables
func createServerInstance() *net.TCPListener {
	serverAddres := fmt.Sprintf("%s:%s", os.Getenv("host"), os.Getenv("port"))
	addr, err := net.ResolveTCPAddr("tcp", serverAddres)
	checkError(err, "cant create serverAddres")

	l, err := net.ListenTCP("tcp", addr)

	checkError(err, "cant create Tcp server")

	return l
}
func printClienInfo(client net.Conn) {
	clientIP, port, _ := net.SplitHostPort(client.RemoteAddr().String())
	fmt.Println("[client connection]")
	fmt.Printf("IP: %s \nport: %s\n", clientIP, port)
	fmt.Println("[end connection]")
}
func main() {
	server := createServerInstance()

	defer server.Close()

	for {
		var data []byte
		var response string
		conn, err := server.Accept()
		checkError(err, "error accepting request")
		printClienInfo(conn)
		n, err := conn.Read(data)
		checkError(err, "error reading client data")
		if n > 0 {
			go handleOperation(data, &response)
		}
		conn.Close()
	}
}
