package main

import (
	"fmt"
	"log"
	"net"
	"os"

	cacheManagement "github.com/EzequielBPullolil/cache_system/cacheManager"
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
		go cacheManagement.HandleOperation(data)
		conn.Close()
	}
}
