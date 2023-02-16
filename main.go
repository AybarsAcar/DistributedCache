package main

import (
	"distributedCache/cache"
	"log"
	"net"
	"time"
)

func main() {
	println("hello, world")

	opts := ServerOptions{
		ListenAddr: ":3000",
		IsLeader:   true,
	}

	// create a client
	go func() {
		time.Sleep(time.Second * 2)

		conn, err := net.Dial("tcp", ":3000")

		if err != nil {
			log.Fatal(err)
		}

		conn.Write([]byte("SET Foo Bar 2500"))
	}()

	server := NewServer(opts, cache.New())

	err := server.Start()
	if err != nil {
		return
	}
}
