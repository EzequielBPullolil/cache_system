package main

import (
	"fmt"
	"net"
	"os"

	handleoperation "github.com/EzequielBPullolil/cache_system/tcp_server/handle_operation"
)

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
		data := make([]byte, 2048)
		response := ""
		conn, err := server.Accept()
		checkError(err, "error accepting request")
		defer conn.Close()
		printClienInfo(conn)
		n, err := conn.Read(data)
		checkError(err, "error reading client data")
		if n > 0 {
			go handleoperation.HandleOperation(data, &response)
		} else {
			response = "operation not found"
		}
		// response to client and close connection
		conn.Write([]byte(response))
		conn.Close()
	}
}
