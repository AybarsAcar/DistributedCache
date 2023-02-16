package main

import (
	"distributedCache/cache"
	"fmt"
	"log"
	"net"
)

type ServerOptions struct {
	ListenAddr string
	IsLeader   bool
}

type Server struct {
	ServerOptions

	cache cache.Cacher
}

func NewServer(ops ServerOptions, c cache.Cacher) *Server {
	return &Server{
		ServerOptions: ops,
		cache:         c,
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenAddr)

	if err != nil {
		return fmt.Errorf("listen error: %s", err)
	}

	log.Printf("Server starting is on port [%s]\n", s.ListenAddr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("accept error: [%s]\n", err)
			continue
		}

		go s.handleConn(conn)
	}
}

// Reads loop
func (s *Server) handleConn(conn net.Conn) {

	defer func() {
		// close the connection at the end of the function
		conn.Close()
	}()

	buf := make([]byte, 2048) // create 2048 byte buffer

	for {
		// read into the buffer
		n, err := conn.Read(buf)

		if err != nil {
			log.Printf("conn read error: %s\n", err)
		}

		message := buf[:n]
		fmt.Println(string(message))
	}
}
