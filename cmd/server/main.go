package main

import (
	"context"
	"errors"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"test-faraway/configs"
)

func main() {

	cfg := configs.Config()

	listener, err := net.Listen("tcp", cfg.ServerHost+":"+cfg.ServerPort)
	if err != nil {
		log.Println("failed to create tcp listener:", err)
		panic(err)
	}

	log.Println("tcp server listening on port:", listener.Addr())
	defer listener.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	connections := make(chan net.Conn)

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				if errors.Is(err, net.ErrClosed) {
					return
				}
				log.Println("error accepting connection:", err)
				continue
			}
			connections <- conn
		}
	}()

	server := Server{}
	for {
		select {
		case c := <-connections:
			go server.Serve(ctx, c)
		case <-sig:
			log.Println("received signal, shutting down")
			cancel()
			return
		}
	}
}
