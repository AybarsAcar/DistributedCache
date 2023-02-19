package main

import (
	"distributedCache/cache"
	"fmt"
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

		conn.Write([]byte("SET Foo Bar 25000000000"))

		time.Sleep(time.Second * 2)
		conn.Write([]byte("GET Foo"))

		buf := make([]byte, 1000)
		n, _ := conn.Read(buf)

		fmt.Println(string(buf[:n]))
	}()

	server := NewServer(opts, cache.New())

	err := server.Start()
	if err != nil {
		return
	}
}
