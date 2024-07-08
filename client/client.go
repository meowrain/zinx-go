package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp4", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	message := "hi there"
	for {
		err = sendMessage(conn, message)
		if err != nil {
			log.Fatalf("Failed to send message: %v", err)
		}
		response, err := readMessage(conn)
		if err != nil {
			log.Fatalf("Failed to read response: %v", err)
		}
		fmt.Println("Received response:", response)
		time.Sleep(time.Second)
	}

}

func sendMessage(conn net.Conn, message string) error {
	_, err := conn.Write([]byte(message))
	if err != nil {
		return fmt.Errorf("failed to send message: %v", err)
	}
	return nil
}

func readMessage(conn net.Conn) (string, error) {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}
	return string(buf[:n]), nil
}
