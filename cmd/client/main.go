package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"test-faraway/configs"
)

func main() {
	cfg := configs.Config()
	conn, err := net.Dial("tcp", cfg.ServerHost+":"+cfg.ServerPort)
	if err != nil {
		log.Fatal("failed to connect to the server", err)
	}

	defer conn.Close()

	bytes, err := io.ReadAll(conn)
	if err != nil {
		log.Fatal("failed to read from the server", err)
	}

	fmt.Println("challenge response:", string(bytes))

}
