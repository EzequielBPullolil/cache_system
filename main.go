package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:29033")

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		fmt.Println("Connection created")
		conn.Close()
	}
}
