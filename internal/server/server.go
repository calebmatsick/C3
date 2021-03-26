package main

import (
	"bufio"
	// "encoding/gob"
	"fmt"
	"net"
	"os"
	"strings"
)

/*
type P struct {
	M, N int64
}
*/

func handleConnection(conn net.Conn) {

	reader := bufio.NewReader(os.Stdin)
	//var cmd string
	fmt.Printf("[shell]: ")
	cmd, _ := reader.ReadString('\n')
	cmd = strings.TrimSuffix(cmd, "\n")
	//encoder := gob.NewEncoder(conn)
	//encoder.Encode(cmd)
	fmt.Fprintf(conn, cmd)
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
