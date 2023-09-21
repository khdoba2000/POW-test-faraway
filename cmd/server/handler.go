package main

import (
	"context"
	"log"
	"net"
)

type Server struct {
}

func (h *Server) Serve(ctx context.Context, conn net.Conn) error {

	defer conn.Close()
	go func() {
		<-ctx.Done()
		conn.Close()
	}()

	log.Println("received conn:", conn)
	conn.Write([]byte("hello from server"))
	return nil
}
