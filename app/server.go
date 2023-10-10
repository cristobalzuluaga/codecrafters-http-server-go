package main

import (
	"fmt"
	"net"
	"os"
	"strings"
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
	n, err := conn.Read(readBuffer)
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
		return
	}

	reqData := string(readBuffer[:n])
	urlPath := parseURLPath(reqData)
	userAgent := parseUserAgent(reqData)

	switch {
	case urlPath == "/":
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	case strings.Contains(urlPath, "/echo/"):
		str := strings.Split(urlPath, "/echo/")
		body := str[1]

		headers := fmt.Sprintf(
			"HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n",
			len(body),
		)

		conn.Write([]byte(headers))
		conn.Write([]byte(body))
	case strings.Contains(urlPath, "/user-agent"):
		body := strings.TrimSpace(userAgent)
		headers := fmt.Sprintf(
			"HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n",
			len(body),
		)

		conn.Write([]byte(headers))
		conn.Write([]byte(body))
	default:
		conn.Write([]byte("HTTP/1.1 404 Not Found response\r\n\r\n"))
	}
}

func parseURLPath(requestData string) string {
	lines := strings.Split(requestData, "\n")
	if len(lines) > 0 {
		parts := strings.Split(lines[0], " ")
		if len(parts) > 1 {
			return parts[1]
		}
	}
	return ""
}

func parseUserAgent(requestData string) string {
	lines := strings.Split(requestData, "\n")
	for _, line := range lines {
		if strings.Contains(line, "User-Agent:") {
			return strings.Split(line, "User-Agent: ")[1]
		}
	}
	return ""
}
