package main

import "time"

type Command uint8

type Message struct {
	Cmd   Command
	Key   []byte
	Value []byte
	TTL   time.Duration
}

const (
	CMDSet Command = iota
	CMDGet
)

func parseCommand(raw []byte) (*Message, error) {

}
