package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/EzequielBPullolil/cache_system/shell/command"
)

func createConnection(host, port string) net.Conn {
	fmt.Println(host + ":" + port)
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		panic(err)
	}

	return conn
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	host := flag.String("host", "", "The host of the tcp connection")
	port := flag.String("port", "", "The port of the tcp connection")
	flag.Parse()
	fmt.Print("\033[H\033[2J")
	fmt.Println("Cache shell")

	for {
		fmt.Print("Cachex> ")
		cmd, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		connection := createConnection(*host, *port)
		command.CommandHandler(cmd, connection)

		connection.Close()
	}

}
