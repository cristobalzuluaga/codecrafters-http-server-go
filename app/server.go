package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer l.Close()

	fmt.Println("Listening on 0.0.0.0:4221")

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	readBuffer := make([]byte, 1024)

	_, err := conn.Read(readBuffer)
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
		return
	}
	fmt.Println("Received: ", string(readBuffer))

	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
}
