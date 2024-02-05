package main

import (
	"fmt"
	"log"
	"net"
	"os"

	handleoperation "github.com/EzequielBPullolil/cache_system/tcp_server/handle_operation"
)

// Stop app if err != nil
func checkError(err error, error_message string) {
	if err != nil {
		log.Fatal(error_message)
		log.Fatal(err)
		os.Exit(1)
	}
}

// Create an TCP server instance using host and port --env variables
func createServerInstance(host, port string) *net.TCPListener {
	serverAddres := host + ":" + port
	addr, err := net.ResolveTCPAddr("tcp", serverAddres)
	checkError(err, "cant create serverAddres")

	l, err := net.ListenTCP("tcp", addr)

	checkError(err, "cant create Tcp server, try using another port")

	return l
}
func printClienInfo(client net.Conn) {
	clientIP, port, _ := net.SplitHostPort(client.RemoteAddr().String())
	log.Printf("[client connection] address: %s:%s\n", clientIP, port)
}

func changeValueOrDefault(slice []string, index int, default_ string) string {
	if len(slice) > index && index >= 0 {
		return slice[index]
	}

	return default_
}
func main() {
	host := changeValueOrDefault(os.Args, 1, "localhost")
	port := changeValueOrDefault(os.Args, 2, "29033")
	server := createServerInstance(host, port)
	fmt.Println("Start Cache_system server at: " + host + ":" + port)
	defer server.Close()

	for {
		data := make([]byte, 1024*3) // Explicit byte
		conn, err := server.Accept()
		checkError(err, "error accepting request")
		defer conn.Close()
		printClienInfo(conn)
		n, err := conn.Read(data)
		checkError(err, "error reading client data")
		if n > 0 {
			go handleoperation.HandleOperation(data) // here change the response variable
		}
		// response to client and close connection
		conn.Write(data)
		log.Println("[end connection]")

		conn.Close()
	}
}
