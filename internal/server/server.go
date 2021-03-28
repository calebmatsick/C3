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
	for {
		fmt.Printf("[shell]: ")
		reader := bufio.NewReader(os.Stdin)

		cmd, _ := reader.ReadString('\n')

		if strings.Compare(cmd, "close\n") == 0 {
			fmt.Fprintf(conn, cmd)
			break
		}

		//encoder := gob.NewEncoder(conn)
		//encoder.Encode(cmd)

		fmt.Fprintf(conn, cmd)

		result, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(string(result))
	}
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
