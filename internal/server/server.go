package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

type P struct {
	M, N int64
}

func handleConnection(conn net.Conn) {
	dec := gob.NewDecoder(conn)
	p := &P{}
	dec.Decode(p)
	fmt.Printf("Recieved : %+v", p)
	conn.Close()
}

func main() {
	fmt.Println("Start server...")

	// Listen on port 8000
	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		// handle error
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			continue
		}
		go handleConnection(conn)
	}
}
