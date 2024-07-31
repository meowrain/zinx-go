package main

import (
	"fmt"
	"log"
	"net"

	"zinx/znet"
)

func main() {
	datapack := znet.NewDataPack()
	msg1 := new(znet.Message)
	msg2 := new(znet.Message)
	str1 := "helloworld,meowrain1"
	str2 := "helloworld,meowrain2"
	msg1.SetMessageLen(uint32(len(str1)))
	msg1.SetMessageId(1)
	msg1.SetData([]byte(str1))
	msg2.SetMessageLen(uint32(len(str2)))
	msg2.SetMessageId(2)
	msg2.SetData([]byte(str2))

	pack1, err := datapack.Pack(msg1)
	if err != nil {
		log.Fatalf("Failed to pack message: %v", err)
	}
	pack2, err := datapack.Pack(msg2)
	if err != nil {
		log.Fatalf("Failed to pack message: %v", err)
	}
	pack := append(pack1, pack2...)
	if err != nil {
		log.Fatalf("Failed to pack message: %v", err)
	}

	conn, err := net.Dial("tcp4", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	_, err = conn.Write(pack)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	// 读取服务器的响应
	response, err := readResponse(conn)
	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
	}
	fmt.Println("Received response:", string(response))
}

func readResponse(conn net.Conn) ([]byte, error) {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}
	return buf[:n], nil
}
