package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide the Host and Port to connect to!!!!!!")
		os.Exit(1)
	}
	host := os.Args[1]
	port := os.Args[2]
	conn, _ := net.Dial("tcp", host+":"+port)
	handle(conn)
}

func handle(conn net.Conn) {
	fmt.Println("Sending: Hello in Go")
	fmt.Fprintf(conn, "Hello in Go\n")
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		conn.Close()
		os.Exit(0)
	}
	fmt.Print("Recieved: " + message)
	conn.Close()
}
