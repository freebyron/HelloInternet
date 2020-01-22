package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a port to listen on!!!")
		os.Exit(1)
	}
	port := os.Args[1]
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("server started at port " + port)
	for i := 0; ; i++ {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		conn.Close()
		os.Exit(0)
	}
	fmt.Print("Recieved: " + message)

	reply := "GoodBye " + strings.Fields(message)[2] + " from Go\n"
	fmt.Println("Sending: " + reply)
	fmt.Fprintf(conn, reply)
	conn.Close()
}
