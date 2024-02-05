package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/EzequielBPullolil/cache_system/shell/command"
)

func createConnection(host, port string) net.Conn {
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		panic(err)
	}

	return conn
}

func changeValueOrDefault(slice []string, index int, default_ string) string {
	if len(slice) > index && index >= 0 {
		return slice[index]
	}

	return default_
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	host := changeValueOrDefault(os.Args, 1, "localhost")
	port := changeValueOrDefault(os.Args, 2, "29033")
	fmt.Print("\033[H\033[2J")
	fmt.Println("Cache shell")

	for {
		fmt.Print("Cachex> ")
		cmd, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		connection := createConnection(host, port)
		defer connection.Close()
		command.CommandHandler(cmd, connection)

	}

}
