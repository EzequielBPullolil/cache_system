package main

import (
	"log"
	"net"
	"os"

	handleoperation "github.com/EzequielBPullolil/cache_system/tcp_server/handle_operation"
)

// Stop app if err != nil
func checkError(err error, error_message string) {
	if err != nil {
		log.Fatal(error_message)
		os.Exit(1)
	}
}

// Create an TCP server instance using host and port --env variables
func createServerInstance() *net.TCPListener {
	serverAddres := "localhost:29033"
	if len(os.Args) > 2 {
		serverAddres = os.Args[1] + ":" + os.Args[2]
	}
	addr, err := net.ResolveTCPAddr("tcp", serverAddres)
	checkError(err, "cant create serverAddres")

	l, err := net.ListenTCP("tcp", addr)

	checkError(err, "cant create Tcp server")

	return l
}
func printClienInfo(client net.Conn) {
	clientIP, port, _ := net.SplitHostPort(client.RemoteAddr().String())
	log.Printf("[client connection] address: %s:%s\n", clientIP, port)
}
func main() {
	server := createServerInstance()

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
