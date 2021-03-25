package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
)

type P struct {
	M, N int64
}

func main() {
	// Start client
	fmt.Println("Start client")
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		log.Fatal("Connection error", err)
	}

	encoder := gob.NewEncoder(conn)
	p := &P{1, 2}
	encoder.Encode(p)
	conn.Close()
	fmt.Println("Finished")
}
