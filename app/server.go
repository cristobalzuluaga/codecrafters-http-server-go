package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	var req []byte

	_, err = conn.Read(req)
	if err != nil {
		log.Println("Error reading bytes: ", err.Error())
		os.Exit(1)
	}

	res, err := conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	if err != nil {
		log.Println("Error writing bytes: ", err.Error())
		os.Exit(1)
	}

	fmt.Printf("res: %v\n", res)
	fmt.Printf("req: %v\n", req)
}
